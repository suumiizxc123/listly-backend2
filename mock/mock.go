package main

import (
	"encoding/json"
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/metal"
	"kcloudb1/internal/utils"
	"log"
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

		tt, _ := time.Parse("2006-01-02", metall.Date)
		met := metal.MetalRate{
			Date:    tt,
			MetalID: 1,
			Rate:    float32(floatValue),
		}
		utils.SendRequestMetal(met)
		time.Sleep(1 * time.Second)

	}
}
