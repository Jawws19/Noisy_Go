package noise

import (
	"fmt"
	"go_noisy_project/pkg/config" // Fix import to the new location
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Define possible protocols and methods
var protocols = []string{"http", "https"}
var methods = []string{"GET", "POST", "PUT", "DELETE", "HEAD"}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/50.0 Safari/537.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Mobile Safari/537.36",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:48.0) Gecko/20100101 Firefox/48.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1 Mobile/15E148 Safari/604.1",
}

// SimulateConcurrentWebTraffic simulates web traffic with custom headers, protocols, and random delays
func SimulateConcurrentWebTraffic(domains []config.Domain, requestCount int, concurrencyLevel int, minDelay int, maxDelay int) {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	requestChannel := make(chan config.Domain)

	// Spawn concurrent workers
	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range requestChannel {
				// Randomly choose a protocol (http or https)
				randomProtocol := protocols[rand.Intn(len(protocols))]

				// Construct the URL using the randomized protocol and domain
				url := fmt.Sprintf("%s://%s", randomProtocol, domain.Domain)

				// Randomly select a User-Agent for each request
				randomUserAgent := userAgents[rand.Intn(len(userAgents))]

				// Randomly choose a method (GET, POST, PUT, DELETE, HEAD)
				randomMethod := methods[rand.Intn(len(methods))]

				// Create the HTTP request with the randomly chosen method
				var req *http.Request
				var err error
				switch randomMethod {
				case "GET":
					req, err = http.NewRequest("GET", url, nil)
				case "POST":
					req, err = http.NewRequest("POST", url, nil) // Add body if needed
				case "PUT":
					req, err = http.NewRequest("PUT", url, nil)
				case "DELETE":
					req, err = http.NewRequest("DELETE", url, nil)
				case "HEAD":
					req, err = http.NewRequest("HEAD", url, nil)
				default:
					fmt.Println("Unsupported method:", randomMethod)
					continue
				}

				if err != nil {
					fmt.Println("Error creating request:", err)
					continue
				}

				// Add custom headers (including the random User-Agent)
				for key, value := range domain.Headers {
					req.Header.Set(key, value)
				}
				// Set the User-Agent header with the randomly selected value
				req.Header.Set("User-Agent", randomUserAgent)

				// Send the request
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}

				fmt.Printf("Request to %s with method %s and User-Agent %s returned status: %s\n", url, randomMethod, randomUserAgent, resp.Status)
				resp.Body.Close()

				// Random delay between minDelay and maxDelay
				randomDelay := time.Duration(rand.Intn(maxDelay-minDelay)+minDelay) * time.Second
				time.Sleep(randomDelay)
			}
		}()
	}

	// Send the requests
	for i := 0; i < requestCount; i++ {
		domain := domains[rand.Intn(len(domains))] // Randomly select a domain
		requestChannel <- domain
	}
	close(requestChannel)

	wg.Wait()
}
