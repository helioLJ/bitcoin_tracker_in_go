package tracker

import "time"

// PriceData represents the Bitcoin price information
type PriceData struct {
	Price     float64
	Currency  string
	Timestamp time.Time
}

// Alert represents a price alert configuration
type Alert struct {
	Type      string // "above" or "below"
	Threshold float64
	Triggered bool
}
