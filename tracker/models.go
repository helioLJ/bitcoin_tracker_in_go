package tracker

import "time"

// PriceData represents the Bitcoin price information
type PriceData struct {
	Price     float64
	Currency  string
	Timestamp time.Time
}

// Config represents the application configuration
type Config struct {
	UpdateInterval time.Duration
	Currency       string
	AlertThreshold float64
	APIKey         string
}

// Alert represents a price alert configuration
type Alert struct {
	Type      string // "above" or "below"
	Threshold float64
	Triggered bool
}
