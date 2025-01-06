package main

import (
	"bitcoin-tracker/tracker"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"bitcoin-tracker/utils"
)

func main() {
	fmt.Println("🪙 Bitcoin Price Tracker")
	fmt.Println("----------------------")
	fmt.Print("Enter currency code (e.g., usd, eur, gbp): ")
	
	var currency string
	fmt.Scanln(&currency)
	
	currency = strings.ToLower(strings.TrimSpace(currency))
	if currency == "" {
		currency = "usd"  // Default to USD if no input
	}

	// Load base config from environment
	config := utils.LoadConfig()
	
	// Override currency if provided via CLI
	if currency != "" {
		config.Currency = currency
	}

	t := tracker.NewTracker(config)
	t.Start()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Main loop
	for {
		select {
		case price := <-t.GetPriceChan():
			fmt.Printf("Bitcoin price: %.2f %s at %s\n", price.Price, price.Currency, price.Timestamp.Format(time.RFC3339))

		case err := <-t.GetErrorChan():
			log.Printf("Error: %v\n", err)

		case <-sigChan:
			fmt.Println("\nShutting down...")
			t.Stop()
			os.Exit(0)
		}
	}
}
