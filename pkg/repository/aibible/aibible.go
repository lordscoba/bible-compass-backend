package gptbible

import (
	"github.com/go-resty/resty/v2"
	"github.com/lordscoba/bible_compass_backend/internal/config"
)

func GptBible(prompt string) (*resty.Response, error) {
	url := config.GetConfig().AiBible.Url

	// Create a new Resty client
	client := resty.New()

	urlMain := url + prompt

	// fmt.Println(urlMain)

	// Send the POST request to Paystack API
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(prompt).
		Get(urlMain)

	if err != nil {
		return nil, err
	}

	// fmt.Println(response.Body())

	return response, nil

}
