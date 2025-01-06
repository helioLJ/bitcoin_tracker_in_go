package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"bitcoin-tracker/utils"
)

const COINGECKO_API_URL = "https://api.coingecko.com/api/v3"

type CoinGeckoResponse struct {
	Bitcoin map[string]float64 `json:"bitcoin"`
}

func FetchBitcoinPrice(currency string) (*PriceData, error) {
	url := fmt.Sprintf("%s/simple/price?ids=bitcoin&vs_currencies=%s", COINGECKO_API_URL, currency)

	resp, err := http.Get(url)
	if err != nil {
		return nil, &utils.CustomError{
			Context: "HTTP request failed",
			Err:     err,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &utils.APIError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}

	var data CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	price, ok := data.Bitcoin[currency]
	if !ok {
		return nil, fmt.Errorf("price not found for currency: %s (API might be rate limited)", currency)
	}

	return &PriceData{
		Price:     price,
		Currency:  currency,
		Timestamp: time.Now(),
	}, nil
}
