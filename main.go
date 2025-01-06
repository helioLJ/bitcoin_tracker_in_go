package main

import (
	"bitcoin-tracker/tracker"
	"bitcoin-tracker/utils"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("🪙 Bitcoin Price Tracker")
	fmt.Println("----------------------")
	fmt.Print("Enter currency code (e.g., usd, eur, gbp): ")

	var currency string
	fmt.Scanln(&currency)

	currency = strings.ToLower(strings.TrimSpace(currency))
	if currency == "" {
		currency = "usd" // Default to USD if no input
	}

	// Load base config from environment
	config := utils.LoadConfig()

	// Override currency if provided via CLI
	if currency != "" {
		config.Currency = currency
	}

	t := tracker.NewTracker(config)

	// Add alerts based on config
	if config.AlertThreshold > 0 {
		t.AddAlert(config.AlertThreshold, "above")
		// Optionally add a below alert at 90% of threshold
		t.AddAlert(config.AlertThreshold*0.9, "below")
	}

	t.Start()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Clear screen and print static UI elements
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println("🪙 Bitcoin Price Tracker")
	fmt.Println("----------------------")
	fmt.Printf("Currency: %s\n", strings.ToUpper(config.Currency))
	fmt.Println("Press Ctrl+C to exit")
	fmt.Println()

	// Main loop
	for {
		select {
		case price := <-t.GetPriceChan():
			sessionChange := t.GetSessionChange(price.Price)
			fmt.Printf("\033[1A\033[2K\rPrice: %.2f %s | Change: %+.2f%% | Updated: %s",
				price.Price,
				strings.ToUpper(price.Currency),
				sessionChange,
				price.Timestamp.Format("15:04:05"))

		case err := <-t.GetErrorChan():
			fmt.Printf("\033[1A\033[2K\rError: %v", err)

		case alert := <-t.GetAlertChan():
			// Print alert on new line, then reprint the price line
			fmt.Printf("\n🚨 Alert: %s\n", alert)
			fmt.Println() // Add extra line for price updates

		case <-sigChan:
			fmt.Println("\nShutting down...")
			t.Stop()
			os.Exit(0)
		}
	}
}
