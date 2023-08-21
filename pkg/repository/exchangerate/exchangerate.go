package raterep

import (
	"github.com/go-resty/resty/v2"
	"github.com/lordscoba/bible_compass_backend/internal/config"
)

func RateRep() (*resty.Response, error) {
	url := config.GetConfig().ExchangeRate.GetUrl

	// Create a new Resty client
	client := resty.New()

	urlMain := url

	// fmt.Println(url)

	// Send the POST request to Paystack API
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(urlMain)

	if err != nil {
		return nil, err
	}

	return response, nil

}
