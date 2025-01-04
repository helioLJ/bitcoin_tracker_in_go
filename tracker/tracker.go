package tracker

import (
	"fmt"
	"time"
)

type Tracker struct {
	config    Config
	alerts    []Alert
	priceChan chan PriceData
	errorChan chan error
	stopChan  chan struct{}
}

func NewTracker(config Config) *Tracker {
	return &Tracker{
		config:    config,
		alerts:    make([]Alert, 0),
		priceChan: make(chan PriceData),
		errorChan: make(chan error),
		stopChan:  make(chan struct{}),
	}
}

func (t *Tracker) Start() {
	go t.trackPrices()
}

func (t *Tracker) Stop() {
	close(t.stopChan)
}

func (t *Tracker) trackPrices() {
	ticker := time.NewTicker(t.config.UpdateInterval)
	defer ticker.Stop()

	fmt.Printf("🔄 Fetching Bitcoin price every %v...\n", t.config.UpdateInterval)
	fmt.Println("Press Ctrl+C to exit")

	// Fetch immediately first
	fmt.Print("Fetching latest price... ")
	if price, err := FetchBitcoinPrice(t.config.Currency); err != nil {
		fmt.Println("❌")
		t.errorChan <- err
	} else {
		fmt.Println("✅")
		t.priceChan <- *price
	}

	for {
		select {
		case <-ticker.C:
			fmt.Print("Fetching latest price... ")
			price, err := FetchBitcoinPrice(t.config.Currency)
			if err != nil {
				fmt.Println("❌")
				t.errorChan <- err
				continue
			}
			fmt.Println("✅")
			t.priceChan <- *price
		case <-t.stopChan:
			return
		}
	}
}

func (t *Tracker) GetPriceChan() <-chan PriceData {
	return t.priceChan
}

func (t *Tracker) GetErrorChan() <-chan error {
	return t.errorChan
}
