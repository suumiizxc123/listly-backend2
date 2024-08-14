package cron

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/models/payment"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func CronJob() {

	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(1).Hour().Do(runJobHour)
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return
	}

	_, err = scheduler.Every(2).Hours().Do(runJobHour2)
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return
	}

	_, err = scheduler.Every(3).Hour().Do(qpayTokenReset)
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return
	}

	// Start the scheduler in the background
	scheduler.StartAsync()

	// Block the main goroutine to keep the program running
	select {}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func qpayTokenReset() {
	config.ConnectDatabase()

	var token payment.QPayToken

	url := "https://merchant.qpay.mn/v2/auth/token"

	// Define your username and password
	username := "LISTLY_AGENT"
	password := "nbIqiJvG"

	// Encode the credentials in Base64
	auth := basicAuth(username, password)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header with the encoded credentials
	req.Header.Set("Authorization", "Basic "+auth)

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status and body
	fmt.Println("QPAY token response Status:", resp.Status)
	if resp.Body != nil {
		// Read response body
		body, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(body, &token)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		config.DB.Create(&token)

		fmt.Println("QPAY token completed")
	}

}

func runJobHour() {
	fmt.Println("Executing cron job at", time.Now().Format(time.RFC3339))
	today := time.Now().Format("2006-01-02")
	prev := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	url := "https://www.mongolbank.mn/mn/gold-and-silver-price/data?startDate=" + prev + "&endDate=" + today

	// Create POST request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response into map[string]interface{}
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// fmt.Println("Response:", jsonResponse["data"])

	dataSlice, ok := jsonResponse["data"].([]interface{})
	if !ok {
		fmt.Println("jsonResponse['data'] is not a []interface{}")
		return
	}

	for i, item := range dataSlice {
		fmt.Printf("Element %d:\n", i+1)

		itemType, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("item is not a map[string]interface{}")
			return
		}

		isInsert := false
		var insertDate time.Time
		var metalRate metal.MetalRate
		for key, value := range itemType {
			if key == "RATE_DATE" {
				fmt.Println("Date: ", value)
				ok, err := metalRate.ExistDate(value.(string))
				if err != nil {
					fmt.Println("Error checking if exists:", err)
					return
				}

				if !ok {
					isInsert = true
					insertDate, _ = time.Parse("2006-01-02", value.(string))
				}

				fmt.Print(" Is insert: ", isInsert)
			}

			if key == "GOLD_BUY" {
				if isInsert {
					metalRate.Date = insertDate
					metalRate.MetalID = 1
					ratevalue, _ := strconv.ParseFloat(strings.ReplaceAll(value.(string), ",", ""), 32)
					metalRate.Rate = float32(ratevalue)

					if err := utils.SendRequestMetal(metalRate); err != nil {
						fmt.Println("Error sending request:", err)
						return
					}
				}
			}
		}
		fmt.Println()
	}
}

func runJobHour2() {
	fmt.Println("Executing cron job at", time.Now().Format(time.RFC3339))

	today := time.Now().Format("2006-01-02")
	prev := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	url := "https://www.mongolbank.mn/mn/currency-rate/data?startDate=" + prev + "&endDate=" + today

	// Create POST request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response into map[string]interface{}
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// fmt.Println("Response:", jsonResponse["data"])

	dataSlice, ok := jsonResponse["data"].([]interface{})
	if !ok {
		fmt.Println("jsonResponse['data'] is not a []interface{}")
		return
	}

	for i, item := range dataSlice {
		fmt.Printf("Element %d:\n", i+1)

		itemType, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("item is not a map[string]interface{}")
			return
		}

		isInsert := false
		var insertDate time.Time
		var metalRate metal.MetalRate
		for key, value := range itemType {
			if key == "RATE_DATE" {
				fmt.Print("Date: ", value)
				ok, err := metalRate.ExistDate(value.(string))
				if err != nil {
					fmt.Println("Error checking if exists:", err)
					return
				}

				if !ok {
					isInsert = true
					insertDate, _ = time.Parse("2006-01-02", value.(string))
				}

				fmt.Print(" Is insert: ", isInsert)
			}

			if key == "USD" {
				if isInsert {
					metalRate.Date = insertDate
					metalRate.MetalID = 2
					ratevalue, _ := strconv.ParseFloat(strings.ReplaceAll(value.(string), ",", ""), 32)
					metalRate.Rate = float32(ratevalue)

					if err := utils.SendCurrencyUSD(metalRate); err != nil {
						fmt.Println("Error sending request:", err)
						return
					}
				}
			}
		}
		fmt.Println()
	}
}
