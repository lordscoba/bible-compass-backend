package randombible

import (
	"github.com/go-resty/resty/v2"
	"github.com/lordscoba/bible_compass_backend/internal/config"
)

func RandomBible() (*resty.Response, error) {
	url := config.GetConfig().RandomBible.Url

	// Create a new Resty client
	client := resty.New()

	// Send the POST request to Paystack API
	response, err := client.R().
		SetHeader("Content-Type", "application/json").Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil

}
