package testdata

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// fixedTime is used for consistent benchmark results.
var fixedTime = time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

// GenerateSmallStruct generates a small structure test data.
func GenerateSmallStruct(seed int) SmallStruct {
	return SmallStruct{
		ID:        int64(seed),
		Name:      fmt.Sprintf("User_%d", seed),
		Email:     fmt.Sprintf("user%d@example.com", seed),
		Age:       20 + (seed % 50),
		Active:    seed%2 == 0,
		Score:     float64(seed) * 1.5,
		Tags:      []string{"developer", "golang", "benchmark"},
		CreatedAt: fixedTime.Format(time.RFC3339),
		UpdatedAt: fixedTime.Add(24 * time.Hour).Format(time.RFC3339),
		Version:   seed % 10,
	}
}

// GenerateLargeStruct generates a large structure test data.
func GenerateLargeStruct(seed int) LargeStruct {
	r := rand.New(rand.NewSource(int64(seed)))

	// Generate 10 order items for consistent benchmark
	itemCount := 10
	items := make([]OrderItem, itemCount)
	for i := range items {
		unitPrice := 10.0 + float64(r.Intn(1000))/10
		quantity := 1 + r.Intn(10)
		discount := float64(r.Intn(20)) / 100
		items[i] = OrderItem{
			ProductID:   int64(1000 + i),
			ProductName: fmt.Sprintf("Product %d", i),
			Quantity:    quantity,
			UnitPrice:   unitPrice,
			Discount:    discount,
			TotalPrice:  float64(quantity) * unitPrice * (1 - discount),
		}
	}

	return LargeStruct{
		OrderID:  fmt.Sprintf("ORD-%d-1704105000000000000", seed),
		UserID:   int64(seed),
		Username: fmt.Sprintf("user_%d", seed),
		Email:    fmt.Sprintf("user%d@example.com", seed),
		Phone:    fmt.Sprintf("+1-555-%04d", seed%10000),
		ShippingAddress: Address{
			Street:     fmt.Sprintf("%d Main St", 100+seed),
			City:       "San Francisco",
			State:      "CA",
			Country:    "USA",
			PostalCode: fmt.Sprintf("%05d", 10000+seed%90000),
			Latitude:   37.7749 + float64(r.Intn(100))/1000,
			Longitude:  -122.4194 + float64(r.Intn(100))/1000,
		},
		BillingAddress: Address{
			Street:     fmt.Sprintf("%d Oak Ave", 200+seed),
			City:       "Los Angeles",
			State:      "CA",
			Country:    "USA",
			PostalCode: fmt.Sprintf("%05d", 90000+seed%10000),
			Latitude:   34.0522 + float64(r.Intn(100))/1000,
			Longitude:  -118.2437 + float64(r.Intn(100))/1000,
		},
		Items: items,
		Payment: PaymentInfo{
			Method:        "credit_card",
			CardLast4:     fmt.Sprintf("%04d", seed%10000),
			TransactionID: fmt.Sprintf("TXN-%d", 1704105000000000000),
			Amount:        1234.56,
			Currency:      "USD",
			Status:        "completed",
			PaidAt:        fixedTime.Format(time.RFC3339),
		},
		Metadata: map[string]string{
			"source":     "web",
			"campaign":   "summer_sale",
			"referrer":   "google",
			"session_id": fmt.Sprintf("sess_%d", seed),
		},
		Status:       "processing",
		Priority:     1 + r.Intn(5),
		Notes:        []string{"Handle with care", "Gift wrap requested"},
		InternalTags: []string{"vip", "recurring", "promo"},
		CreatedAt:    fixedTime.Add(-24 * time.Hour).Format(time.RFC3339),
		UpdatedAt:    fixedTime.Format(time.RFC3339),
		SubTotal:     1000.00,
		TaxAmount:    80.00,
		ShippingCost: 15.00,
		TotalAmount:  1095.00,
		IsGift:       seed%3 == 0,
		RequireSign:  true,
		ExpressShip:  seed%2 == 0,
		Verified:     true,
	}
}

// GenerateSmallStructJSON generates JSON bytes for small structure.
func GenerateSmallStructJSON(seed int) []byte {
	s := GenerateSmallStruct(seed)
	data, _ := json.Marshal(s)
	return data
}

// GenerateLargeStructJSON generates JSON bytes for large structure.
func GenerateLargeStructJSON(seed int) []byte {
	s := GenerateLargeStruct(seed)
	data, _ := json.Marshal(s)
	return data
}
