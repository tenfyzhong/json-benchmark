//go:generate easyjson -all types.go

// Package testdata provides test data structures for JSON benchmark.
package testdata

// SmallStruct represents a simple small structure, simulating common API responses.
// About 10 fields with shallow nesting.
type SmallStruct struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Active    bool     `json:"active"`
	Score     float64  `json:"score"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Version   int      `json:"version"`
}

// Address represents address information.
type Address struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

// OrderItem represents an order line item.
type OrderItem struct {
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Discount    float64 `json:"discount"`
	TotalPrice  float64 `json:"total_price"`
}

// PaymentInfo represents payment information.
type PaymentInfo struct {
	Method        string  `json:"method"`
	CardLast4     string  `json:"card_last4,omitempty"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
	PaidAt        string  `json:"paid_at"`
}

// LargeStruct represents a complex large structure, simulating an e-commerce order.
// Contains nested structures, arrays, and maps, about 50+ fields.
type LargeStruct struct {
	// Basic information
	OrderID  string `json:"order_id"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	// Address information
	ShippingAddress Address `json:"shipping_address"`
	BillingAddress  Address `json:"billing_address"`

	// Order items - array
	Items []OrderItem `json:"items"`

	// Payment information
	Payment PaymentInfo `json:"payment"`

	// Metadata - map
	Metadata map[string]string `json:"metadata"`

	// Status information
	Status       string   `json:"status"`
	Priority     int      `json:"priority"`
	Notes        []string `json:"notes"`
	InternalTags []string `json:"internal_tags"`

	// Timestamps
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CompletedAt string `json:"completed_at,omitempty"`

	// Statistics
	SubTotal     float64 `json:"sub_total"`
	TaxAmount    float64 `json:"tax_amount"`
	ShippingCost float64 `json:"shipping_cost"`
	TotalAmount  float64 `json:"total_amount"`

	// Boolean flags
	IsGift      bool `json:"is_gift"`
	RequireSign bool `json:"require_sign"`
	ExpressShip bool `json:"express_ship"`
	Verified    bool `json:"verified"`
}
