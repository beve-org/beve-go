package core

import (
    "runtime"
    "sync"
    "testing"
)

// BenchmarkPoolContention measures encoder pool overhead.

func BenchmarkPoolSerial(b *testing.B) {
    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        enc := GetEncoderFromPool()
        PutEncoderToPool(enc)
    }
}

func BenchmarkPoolParallel(b *testing.B) {
    b.Run("4-goroutines", func(b *testing.B) {
        b.ReportAllocs()
        b.SetParallelism(4)
        b.ResetTimer()

        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                enc := GetEncoderFromPool()
                PutEncoderToPool(enc)
            }
        })
    })

    b.Run("8-goroutines", func(b *testing.B) {
        b.ReportAllocs()
        b.SetParallelism(8)
        b.ResetTimer()

        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                enc := GetEncoderFromPool()
                PutEncoderToPool(enc)
            }
        })
    })

    b.Run("16-goroutines", func(b *testing.B) {
        b.ReportAllocs()
        b.SetParallelism(16)
        b.ResetTimer()

        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                enc := GetEncoderFromPool()
                PutEncoderToPool(enc)
            }
        })
    })
}

func BenchmarkPoolVsAlloc(b *testing.B) {
    b.Run("Pool", func(b *testing.B) {
        b.ReportAllocs()
        b.ResetTimer()

        for i := 0; i < b.N; i++ {
            enc := GetEncoderFromPool()
            PutEncoderToPool(enc)
        }
    })

    b.Run("Alloc", func(b *testing.B) {
        b.ReportAllocs()
        b.ResetTimer()

        for i := 0; i < b.N; i++ {
            buf := AcquireBuffer(1024)
            _ = NewEncoder(buf)
            ReleaseBuffer(buf)
        }
    })
}

func BenchmarkPoolPerCPU(b *testing.B) {
    numCPUs := runtime.NumCPU()
    cpuPools := make([]*sync.Pool, numCPUs)
    
    for i := 0; i < numCPUs; i++ {
        cpuPools[i] = &sync.Pool{
            New: func() interface{} {
                buf := AcquireBuffer(1024)
                return NewEncoder(buf)
            },
        }
    }

    b.Run("GlobalPool/8-goroutines", func(b *testing.B) {
        b.ReportAllocs()
        b.SetParallelism(8)
        b.ResetTimer()

        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                enc := GetEncoderFromPool()
                PutEncoderToPool(enc)
            }
        })
    })

    b.Run("PerCPUPool/8-goroutines", func(b *testing.B) {
        b.ReportAllocs()
        b.SetParallelism(8)
        b.ResetTimer()

        var counter int64
        b.RunParallel(func(pb *testing.PB) {
            cpuID := int(counter) % numCPUs
            counter++
            pool := cpuPools[cpuID]

            for pb.Next() {
                enc := pool.Get().(*Encoder)
                pool.Put(enc)
            }
        })
    })
}
