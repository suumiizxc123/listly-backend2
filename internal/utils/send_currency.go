package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"kcloudb1/internal/models/metal"
	"net/http"
)

func SendCurrencyUSD(mmm metal.MetalRate) error {
	url := "http://oggbackend.999.mn:8080/api/v1/metal-rate"

	// JSON payload to send in the request body
	jsonData := fmt.Sprintf(`{"date": "%v", "metal_id": 2, "rate": %.2f}`, mmm.Date.Format("2006-01-02T15:04:05Z"), mmm.Rate)

	// Create a new request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Make the POST request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Print the response body
	fmt.Println(string(body))

	return nil
}
