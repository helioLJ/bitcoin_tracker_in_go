package tracker

import (
	"fmt"
	"time"
	"bitcoin-tracker/utils"
)

type Tracker struct {
	config    utils.Config
	alerts    []Alert
	priceChan chan PriceData
	errorChan chan error
	stopChan  chan struct{}
	initialPrice float64
	priceHistory []PriceData
	alertChan chan string
}

func NewTracker(config utils.Config) *Tracker {
	return &Tracker{
		config:    config,
		alerts:    make([]Alert, 0),
		priceChan: make(chan PriceData),
		errorChan: make(chan error),
		stopChan:  make(chan struct{}),
		initialPrice: 0,
		priceHistory: make([]PriceData, 0),
		alertChan: make(chan string),
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

	// Fetch initial price
	fmt.Print("Fetching latest price... ")
	if price, err := FetchBitcoinPrice(t.config.Currency); err != nil {
		fmt.Println("❌")
		t.errorChan <- err
	} else {
		fmt.Println("✅")
		t.initialPrice = price.Price
		t.priceHistory = append(t.priceHistory, *price)
		t.priceChan <- *price
		t.logPrice(*price)
	}

	backoff := time.Second
	maxBackoff := time.Minute

	for {
		select {
		case <-ticker.C:
			fmt.Print("\nFetching latest price... ")
			price, err := FetchBitcoinPrice(t.config.Currency)
			if err != nil {
				switch e := err.(type) {
				case *utils.APIError:
					if e.StatusCode == 429 { // Rate limit
						time.Sleep(backoff)
						backoff *= 2
						if backoff > maxBackoff {
							backoff = maxBackoff
						}
					}
				}
				t.errorChan <- err
				continue
			}
			backoff = time.Second // Reset backoff on success
			fmt.Println("✅")
			
			t.priceHistory = append(t.priceHistory, *price)
			t.checkAlerts(price.Price)
			t.priceChan <- *price
			t.logPrice(*price)
			
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

func (t *Tracker) GetSessionChange(currentPrice float64) float64 {
	if t.initialPrice == 0 {
		return 0
	}
	return ((currentPrice - t.initialPrice) / t.initialPrice) * 100
}

func (t *Tracker) logPrice(price PriceData) {
	logEntry := fmt.Sprintf("%s,%.2f,%s\n",
		price.Timestamp.Format(time.RFC3339),
		price.Price,
		price.Currency)
	
	utils.AppendToFile("price_history.csv", logEntry)
}

func (t *Tracker) AddAlert(threshold float64, alertType string) {
	t.alerts = append(t.alerts, Alert{
		Threshold: threshold,
		Type:      alertType,
		Triggered: false,
	})
}

func (t *Tracker) checkAlerts(price float64) {
	for i := range t.alerts {
		alert := &t.alerts[i]
		if alert.Type == "above" && price > alert.Threshold && !alert.Triggered {
			t.alertChan <- fmt.Sprintf("Price above %.2f %s", alert.Threshold, t.config.Currency)
			alert.Triggered = true
		} else if alert.Type == "below" && price < alert.Threshold && !alert.Triggered {
			t.alertChan <- fmt.Sprintf("Price below %.2f %s", alert.Threshold, t.config.Currency)
			alert.Triggered = true
		}
	}
}

func (t *Tracker) GetAlertChan() <-chan string {
	return t.alertChan
}
