package main

import (
	"encoding/json"
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/metal"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	config.ConnectDatabase()
	// Create a new scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	// Schedule a job to run every minute
	_, err := scheduler.Every(1).Minute().Do(func() {
		fmt.Println("Executing cron job at", time.Now().Format(time.RFC3339))
		today := time.Now().Format("2006-01-02")
		prev := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
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

		var metalRate metal.MetalRate
		for i, item := range dataSlice {
			fmt.Printf("Element %d:\n", i+1)

			itemType, ok := item.(map[string]interface{})
			if !ok {
				fmt.Println("item is not a map[string]interface{}")
				return
			}

			isInsert := false
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
					}

					fmt.Println("Is insert: ", isInsert)
				}

			}
			fmt.Println() // Print a newline for separation
		}

	})
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return
	}

	// Start the scheduler (it runs jobs in its own goroutine)
	scheduler.StartAsync()

	// Keep the program running
	select {}
}
