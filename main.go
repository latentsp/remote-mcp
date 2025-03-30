package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type MCPRequest struct {
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Version string          `json:"jsonrpc"`
}

type MCPResponse struct {
	ID      json.RawMessage `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *MCPError       `json:"error,omitempty"`
	Version string          `json:"jsonrpc"`
}

type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	// Define command line flags
	initialEndpointURL := flag.String("endpoint-url", "", "The URL of the remote MCP server")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	// Setup logging to file
	logFile, err := os.OpenFile("/Users/gmaliar/src/latentsp/remote-mcp/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix(fmt.Sprintf("[%s] [remote-mcp] ", time.Now().Format("2006-01-02 15:04:05")))

	// Check if endpoint URL was provided
	if *initialEndpointURL == "" {
		log.Println("Error: --endpoint-url flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if *verbose {
		log.Println("Verbose logging enabled")
		log.Printf("Remote MCP server URL: %s", *initialEndpointURL)
	}

	// Create HTTP client
	client := &http.Client{}

	// Connect to SSE endpoint
	sseURL := *initialEndpointURL + "/sse"
	req, err := http.NewRequest("GET", sseURL, nil)
	if err != nil {
		log.Fatalf("Error creating SSE request: %v", err)
	}
	req.Header.Set("Accept", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error connecting to SSE endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Create channel for responses
	responseChan := make(chan string)
	defer close(responseChan)

	// Create channel to signal when endpoint URL is received
	endpointChan := make(chan string)
	defer close(endpointChan)

	// Start goroutine to read SSE responses
	go func() {
		scanner := bufio.NewScanner(resp.Body)
		var currentEvent string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			// Handle event line
			if strings.HasPrefix(line, "event: ") {
				currentEvent = line[7:] // Remove "event: " prefix
				log.Printf("Received SSE event: %s", currentEvent)
				continue
			}

			if !strings.HasPrefix(line, "data: ") {
				log.Printf("Invalid SSE data format: %s", line)
				continue
			}

			data := line[6:]

			log.Printf("Received SSE data (event: %s): %s", currentEvent, data)

			if currentEvent == "endpoint" {
				log.Printf("Got endpoint URL: %s", data)
				endpointChan <- data
				continue
			}

			if currentEvent == "ping" {
				log.Printf("Skipping ping event")
				continue
			}

			fmt.Fprintln(os.Stdout, data)
		}
	}()

	// Wait for endpoint URL before proceeding
	log.Println("Waiting for endpoint URL...")
	mcpURL := <-endpointChan
	mcpURL = *initialEndpointURL + mcpURL
	log.Printf("Connection established, endpoint URL: %s", mcpURL)

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		line := stdinScanner.Text()
		if line == "" {
			continue
		}

		log.Printf("Received from stdin: %s", line)

		var req MCPRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			log.Printf("Error parsing MCP request: %v", err)
			continue
		}

		httpReq, err := http.NewRequest("POST", mcpURL, bytes.NewBuffer([]byte(line)))
		if err != nil {
			log.Printf("Error creating HTTP request: %v", err)
			continue
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(httpReq)
		if err != nil {
			log.Printf("Error sending request to MCP server: %v", err)
			continue
		}
		resp.Body.Close()
	}

	if err := stdinScanner.Err(); err != nil {
		log.Fatalf("Error reading stdin: %v", err)
	}
}