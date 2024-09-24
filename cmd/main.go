package main

import (
	"flag"
	"fmt"
	"go_noisy_project/pkg/config" // Updated import path for config
	"go_noisy_project/pkg/noise"  // Updated import path for noise
	"os"
	"strconv"
)

func main() {
	// Get the path to the YAML file from a command-line flag
	yamlFilePath := flag.String("config", "domains.yaml", "Path to the YAML file containing domains")
	flag.Parse()

	// Check if the file exists
	if _, err := os.Stat(*yamlFilePath); os.IsNotExist(err) {
		fmt.Println("YAML file not found:", *yamlFilePath)
		return
	}

	// Load the list of domains from the YAML file
	domains, err := config.LoadDomainsFromYAML(*yamlFilePath)
	if err != nil {
		fmt.Println("Error loading domains:", err)
		return
	}

	fmt.Println("Loaded", len(domains), "domains from the YAML file.")

	// Prompt user for minimum and maximum delays between site visits
	var inputMinDelay, inputMaxDelay string
	fmt.Print("Enter minimum delay between site visits (in seconds): ")
	fmt.Scanln(&inputMinDelay)
	fmt.Print("Enter maximum delay between site visits (in seconds): ")
	fmt.Scanln(&inputMaxDelay)

	// Convert user input to integers
	minDelay, err := strconv.Atoi(inputMinDelay)
	if err != nil {
		fmt.Println("Invalid input for minimum delay.")
		return
	}
	maxDelay, err := strconv.Atoi(inputMaxDelay)
	if err != nil {
		fmt.Println("Invalid input for maximum delay.")
		return
	}

	// Simulate web traffic concurrently with random delays and dynamic headers
	requestCount := 100    // Total number of requests to simulate
	concurrencyLevel := 10 // Number of concurrent goroutines

	fmt.Printf("Starting web traffic simulation with concurrency... Delay: %d-%d seconds\n", minDelay, maxDelay)
	noise.SimulateConcurrentWebTraffic(domains, requestCount, concurrencyLevel, minDelay, maxDelay)

	fmt.Println("Web traffic simulation complete.")
}
