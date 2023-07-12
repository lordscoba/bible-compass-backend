package paystack

import (
	"github.com/go-resty/resty/v2"
	"github.com/lordscoba/bible_compass_backend/internal/config"
)

func PaystackInitPost(payload map[string]interface{}) (*resty.Response, error) {

	// Set up the Paystack API endpoint
	url := config.GetConfig().Paystack.InitUrl
	bearer := config.GetConfig().Paystack.PaystackKey

	// Create a new Resty client
	client := resty.New()

	// Set the Paystack API key in the request header
	client.SetHeader("Authorization", "Bearer "+bearer)

	// Send the POST request to Paystack API
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func PaystackVerifyGet(reference string) (*resty.Response, error) {

	// Set up the Paystack API endpoint
	url := config.GetConfig().Paystack.VerifyUrl
	bearer := config.GetConfig().Paystack.PaystackKey

	// Create a new Resty client
	client := resty.New()

	// Set the Paystack API key in the request header
	client.SetHeader("Authorization", "Bearer "+bearer)

	// Send the POST request to Paystack API
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(url + reference)

	if err != nil {
		return nil, err
	}

	return response, nil
}
