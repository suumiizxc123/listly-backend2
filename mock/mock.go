package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/metal"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BomMetal struct {
	Date   string `json:"date"`
	Gold   string `json:"gold"`
	Silver string `json:"silver"`
}

func main() {
	config.ConnectDatabase()
	getMetal()
}

func getMetal() {
	// Print the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	fmt.Println("Current Directory:", currentDir)

	// Use a relative path to locate the JSON file
	filePath := "goldprev.json"

	// Attempt to read the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", filePath, err)
	}

	// Parse JSON
	var metals []BomMetal
	if err := json.Unmarshal(fileContent, &metals); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Sort metals by date in ascending order
	sort.Slice(metals, func(i, j int) bool {
		date1, _ := time.Parse("2006-01-02", metals[i].Date)
		date2, _ := time.Parse("2006-01-02", metals[j].Date)
		return date1.Before(date2)
	})

	for _, metall := range metals {
		floatValue, err := strconv.ParseFloat(strings.ReplaceAll(metall.Gold, ",", ""), 32)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		tt, err := time.Parse("2006-01-02", metall.Date)
		met := metal.MetalRate{
			Date:    tt,
			MetalID: 1,
			Rate:    float32(floatValue),
		}
		sendRequest(met)
		time.Sleep(1 * time.Second)

	}
}

func sendRequest(mmm metal.MetalRate) error {
	url := "http://localhost:8080/api/v1/metal-rate"

	// JSON payload to send in the request body
	jsonData := fmt.Sprintf(`{"date": "%v", "metal_id": 1, "rate": %.2f}`, mmm.Date.Format("2006-01-02T15:04:05Z"), mmm.Rate)

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
