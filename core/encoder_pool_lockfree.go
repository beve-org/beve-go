//go:build go1.21
// +build go1.21

package core

import (
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

// UseLockFreePool controls whether to use lock-free per-P pools (default: false)
// Can be enabled via BEVE_USE_LOCKFREE_POOL=true environment variable
var UseLockFreePool = false

// lockFreePoolMaxDepth is the maximum number of encoders per P pool
// Prevents unbounded memory growth
const lockFreePoolMaxDepth = 32

// lockFreePoolStats tracks pool performance metrics
type lockFreePoolStats struct {
	hits      uint64 // Successful pool retrievals
	misses    uint64 // Pool empty, created new encoder
	puts      uint64 // Successful pool returns
	discards  uint64 // Encoders discarded (pool full or too large)
	overflows uint64 // Pool depth exceeded maxDepth
}

var globalLockFreeStats lockFreePoolStats

// encoderStack represents a lock-free stack of encoders for a single P
// Cache-line padded to prevent false sharing (128 bytes on ARM64, 64 on AMD64)
type encoderStack struct {
	_     [128]byte   // Leading padding
	head  *Encoder    // Stack head (lock-free linked list)
	depth int32       // Current pool depth (atomic)
	_     [128]byte   // Trailing padding
}

// perPEncoderPools contains one encoderStack per P (CPU core)
// Initialized lazily on first use
var (
	perPEncoderPools     []*encoderStack
	perPEncoderPoolsOnce sync.Once
)

// initPerPPools initializes per-P encoder pools
func initPerPPools() {
	numP := runtime.GOMAXPROCS(0)
	perPEncoderPools = make([]*encoderStack, numP)
	
	for i := 0; i < numP; i++ {
		perPEncoderPools[i] = &encoderStack{}
	}
}

// GetLockFreePoolStats returns current pool statistics
func GetLockFreePoolStats() (hits, misses, puts, discards, overflows uint64) {
	return atomic.LoadUint64(&globalLockFreeStats.hits),
		atomic.LoadUint64(&globalLockFreeStats.misses),
		atomic.LoadUint64(&globalLockFreeStats.puts),
		atomic.LoadUint64(&globalLockFreeStats.discards),
		atomic.LoadUint64(&globalLockFreeStats.overflows)
}

// ResetLockFreePoolStats resets pool statistics (for testing)
func ResetLockFreePoolStats() {
	atomic.StoreUint64(&globalLockFreeStats.hits, 0)
	atomic.StoreUint64(&globalLockFreeStats.misses, 0)
	atomic.StoreUint64(&globalLockFreeStats.puts, 0)
	atomic.StoreUint64(&globalLockFreeStats.discards, 0)
	atomic.StoreUint64(&globalLockFreeStats.overflows, 0)
}

// runtime_procPin prevents the current goroutine from being preempted
// and returns the current P's ID.
//
//go:linkname runtime_procPin runtime.procPin
//go:noescape
func runtime_procPin() int

// runtime_procUnpin allows the current goroutine to be preempted again.
//
//go:linkname runtime_procUnpin runtime.procUnpin
//go:noescape
func runtime_procUnpin()

// getEncoderFromLockFreePool retrieves an encoder from the per-P lock-free pool
func getEncoderFromLockFreePool() *Encoder {
	// Ensure pools are initialized
	perPEncoderPoolsOnce.Do(initPerPPools)
	
	// Pin to current P to prevent migration
	pid := runtime_procPin()
	
	// Safety check: pid should be in valid range
	if pid < 0 || pid >= len(perPEncoderPools) {
		runtime_procUnpin()
		// Fallback to creating new encoder
		atomic.AddUint64(&globalLockFreeStats.misses, 1)
		return NewEncoder(nil)
	}
	
	stack := perPEncoderPools[pid]
	runtime_procUnpin() // Unpin as we got our stack pointer
	
	// Lock-free pop from stack using atomic CAS
	for {
		head := (*Encoder)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&stack.head))))
		if head == nil {
			// Pool is empty, create new encoder
			atomic.AddUint64(&globalLockFreeStats.misses, 1)
			enc := NewEncoder(nil)
			enc.Buf = AcquireBuffer(getOptimalBufferCapacity())
			return enc
		}
		
		// Load next pointer atomically
		next := (*Encoder)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&head.next))))
		
		// Try to atomically swap head with head.next
		// This is the lock-free "pop" operation
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&stack.head)),
			unsafe.Pointer(head),
			unsafe.Pointer(next),
		) {
			// Success! We popped the encoder
			atomic.AddInt32(&stack.depth, -1)
			atomic.AddUint64(&globalLockFreeStats.hits, 1)
			
			// Clear the next pointer (encoder is no longer in list)
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&head.next)), nil)
			return head
		}
		
		// CAS failed (another goroutine modified head), retry
		// Hardware will handle contention with exponential backoff
	}
}

// putEncoderToLockFreePool returns an encoder to the per-P lock-free pool
func putEncoderToLockFreePool(enc *Encoder) {
	if enc == nil || enc.Buf == nil {
		return
	}
	
	// Check buffer size: don't pool overly large buffers
	bufCap := cap(enc.Buf.data)
	if bufCap > maxBufferPoolCapacity {
		atomic.AddUint64(&globalLockFreeStats.discards, 1)
		ReleaseBuffer(enc.Buf)
		return
	}
	
	// Reset encoder state
	enc.Buf.Reset()
	enc.batchLen = 0
	enc.w = nil
	
	// Ensure pools are initialized
	perPEncoderPoolsOnce.Do(initPerPPools)
	
	// Pin to current P
	pid := runtime_procPin()
	
	// Safety check
	if pid < 0 || pid >= len(perPEncoderPools) {
		runtime_procUnpin()
		atomic.AddUint64(&globalLockFreeStats.discards, 1)
		ReleaseBuffer(enc.Buf)
		return
	}
	
	stack := perPEncoderPools[pid]
	runtime_procUnpin()
	
	// Check pool depth to prevent unbounded growth
	currentDepth := atomic.LoadInt32(&stack.depth)
	if currentDepth >= lockFreePoolMaxDepth {
		atomic.AddUint64(&globalLockFreeStats.overflows, 1)
		ReleaseBuffer(enc.Buf)
		return
	}
	
	// Lock-free push to stack using atomic CAS
	for {
		oldHead := (*Encoder)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&stack.head))))
		
		// Link encoder to current head (atomically)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&enc.next)), unsafe.Pointer(oldHead))
		
		// Try to atomically swap head with enc
		// This is the lock-free "push" operation
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&stack.head)),
			unsafe.Pointer(oldHead),
			unsafe.Pointer(enc),
		) {
			// Success! Encoder is now in the pool
			atomic.AddInt32(&stack.depth, 1)
			atomic.AddUint64(&globalLockFreeStats.puts, 1)
			return
		}
		
		// CAS failed, retry
	}
}

func init() {
	// Check environment variable to enable lock-free pool
	if val := os.Getenv("BEVE_USE_LOCKFREE_POOL"); val != "" {
		if enabled, err := strconv.ParseBool(val); err == nil {
			UseLockFreePool = enabled
		}
	}
}
