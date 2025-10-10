package beve

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/fxamacker/cbor/v2"
	"github.com/go-faker/faker/v4"
	"github.com/vmihailenco/msgpack/v5"
)

// Realistic data structures using faker

type User struct {
	ID        int      `json:"id" msgpack:"id" cbor:"id" beve:"id"`
	FirstName string   `json:"first_name" msgpack:"first_name" cbor:"first_name" beve:"first_name" faker:"first_name"`
	LastName  string   `json:"last_name" msgpack:"last_name" cbor:"last_name" beve:"last_name" faker:"last_name"`
	Email     string   `json:"email" msgpack:"email" cbor:"email" beve:"email" faker:"email"`
	Username  string   `json:"username" msgpack:"username" cbor:"username" beve:"username" faker:"username"`
	Phone     string   `json:"phone" msgpack:"phone" cbor:"phone" beve:"phone" faker:"phone_number"`
	Age       int      `json:"age" msgpack:"age" cbor:"age" beve:"age"`
	Balance   float64  `json:"balance" msgpack:"balance" cbor:"balance" beve:"balance"`
	IsActive  bool     `json:"is_active" msgpack:"is_active" cbor:"is_active" beve:"is_active"`
	Tags      []string `json:"tags" msgpack:"tags" cbor:"tags" beve:"tags"`
}

type Order struct {
	OrderID    string  `json:"order_id" msgpack:"order_id" cbor:"order_id" beve:"order_id" faker:"uuid_hyphenated"`
	UserID     int     `json:"user_id" msgpack:"user_id" cbor:"user_id" beve:"user_id"`
	ProductID  string  `json:"product_id" msgpack:"product_id" cbor:"product_id" beve:"product_id" faker:"uuid_hyphenated"`
	Quantity   int     `json:"quantity" msgpack:"quantity" cbor:"quantity" beve:"quantity"`
	Price      float64 `json:"price" msgpack:"price" cbor:"price" beve:"price"`
	TotalPrice float64 `json:"total_price" msgpack:"total_price" cbor:"total_price" beve:"total_price"`
	Status     string  `json:"status" msgpack:"status" cbor:"status" beve:"status"`
	CreatedAt  string  `json:"created_at" msgpack:"created_at" cbor:"created_at" beve:"created_at" faker:"timestamp"`
}

type ComplexData struct {
	Users    []User                 `json:"users" msgpack:"users" cbor:"users" beve:"users"`
	Orders   []Order                `json:"orders" msgpack:"orders" cbor:"orders" beve:"orders"`
	Metadata map[string]interface{} `json:"metadata" msgpack:"metadata" cbor:"metadata" beve:"metadata"`
}

// Helper to generate fake data
func generateUser() User {
	user := User{
		ID:       int(faker.RandomUnixTime()),
		Age:      25,
		Balance:  12345.67,
		IsActive: true,
		Tags:     []string{"premium", "verified", "active"},
	}
	_ = faker.FakeData(&user)
	return user
}

func generateOrder() Order {
	order := Order{
		UserID:     123,
		Quantity:   5,
		Price:      99.99,
		TotalPrice: 499.95,
		Status:     "completed",
	}
	_ = faker.FakeData(&order)
	return order
}

func generateComplexData(userCount, orderCount int) ComplexData {
	data := ComplexData{
		Users:  make([]User, userCount),
		Orders: make([]Order, orderCount),
		Metadata: map[string]interface{}{
			"version":   "1.0",
			"timestamp": faker.Timestamp(),
			"server":    "prod-01",
		},
	}

	for i := 0; i < userCount; i++ {
		data.Users[i] = generateUser()
	}

	for i := 0; i < orderCount; i++ {
		data.Orders[i] = generateOrder()
	}

	return data
}

// ==================== SMALL STRUCT BENCHMARKS ====================

func BenchmarkSmallStruct_BEVE_Marshal(b *testing.B) {
	user := generateUser()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := Marshal(user)
		benchBytesSink = data
	}
}

func BenchmarkSmallStruct_JSON_Marshal(b *testing.B) {
	user := generateUser()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(user)
		benchBytesSink = data
	}
}

func BenchmarkSmallStruct_Sonic_Marshal(b *testing.B) {
	user := generateUser()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := sonic.Marshal(user)
		benchBytesSink = data
	}
}

func BenchmarkSmallStruct_MessagePack_Marshal(b *testing.B) {
	user := generateUser()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := msgpack.Marshal(user)
		benchBytesSink = data
	}
}

func BenchmarkSmallStruct_CBOR_Marshal(b *testing.B) {
	user := generateUser()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, _ := cbor.Marshal(user)
		benchBytesSink = data
	}
}

// ==================== SMALL STRUCT UNMARSHAL ====================

func BenchmarkSmallStruct_BEVE_Unmarshal(b *testing.B) {
	user := generateUser()
	data, _ := Marshal(user)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result User
		_ = Unmarshal(data, &result)
	}
}

func BenchmarkSmallStruct_JSON_Unmarshal(b *testing.B) {
	user := generateUser()
	data, _ := json.Marshal(user)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result User
		_ = json.Unmarshal(data, &result)
	}
}

func BenchmarkSmallStruct_Sonic_Unmarshal(b *testing.B) {
	user := generateUser()
	data, _ := sonic.Marshal(user)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result User
		_ = sonic.Unmarshal(data, &result)
	}
}

func BenchmarkSmallStruct_MessagePack_Unmarshal(b *testing.B) {
	user := generateUser()
	data, _ := msgpack.Marshal(user)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result User
		_ = msgpack.Unmarshal(data, &result)
	}
}

func BenchmarkSmallStruct_CBOR_Unmarshal(b *testing.B) {
	user := generateUser()
	data, _ := cbor.Marshal(user)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var result User
		_ = cbor.Unmarshal(data, &result)
	}
}

// ==================== MEDIUM COMPLEXITY BENCHMARKS ====================

func BenchmarkMedium_BEVE_Marshal(b *testing.B) {
	data := generateComplexData(10, 20)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkMedium_JSON_Marshal(b *testing.B) {
	data := generateComplexData(10, 20)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := json.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkMedium_Sonic_Marshal(b *testing.B) {
	data := generateComplexData(10, 20)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := sonic.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkMedium_MessagePack_Marshal(b *testing.B) {
	data := generateComplexData(10, 20)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := msgpack.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkMedium_CBOR_Marshal(b *testing.B) {
	data := generateComplexData(10, 20)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := cbor.Marshal(data)
		benchBytesSink = result
	}
}

// ==================== LARGE PAYLOAD BENCHMARKS ====================

func BenchmarkLarge_BEVE_Marshal(b *testing.B) {
	data := generateComplexData(100, 200)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkLarge_JSON_Marshal(b *testing.B) {
	data := generateComplexData(100, 200)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := json.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkLarge_Sonic_Marshal(b *testing.B) {
	data := generateComplexData(100, 200)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := sonic.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkLarge_MessagePack_Marshal(b *testing.B) {
	data := generateComplexData(100, 200)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := msgpack.Marshal(data)
		benchBytesSink = result
	}
}

func BenchmarkLarge_CBOR_Marshal(b *testing.B) {
	data := generateComplexData(100, 200)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, _ := cbor.Marshal(data)
		benchBytesSink = result
	}
}

// ==================== SIZE COMPARISON ====================

func BenchmarkSize_BEVE(b *testing.B) {
	user := generateUser()
	data, _ := Marshal(user)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}

func BenchmarkSize_JSON(b *testing.B) {
	user := generateUser()
	data, _ := json.Marshal(user)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}

func BenchmarkSize_Sonic(b *testing.B) {
	user := generateUser()
	data, _ := sonic.Marshal(user)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}

func BenchmarkSize_MessagePack(b *testing.B) {
	user := generateUser()
	data, _ := msgpack.Marshal(user)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}

func BenchmarkSize_CBOR(b *testing.B) {
	user := generateUser()
	data, _ := cbor.Marshal(user)
	b.ReportMetric(float64(len(data)), "bytes")
	b.ReportAllocs()
}
