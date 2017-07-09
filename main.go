package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Indiegogo gets data from Indiegogo
type Indiegogo struct {
	Response []struct {
		By string `json:"by"`
	} `json:"response"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//handler echoes the donors
func handler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://api.indiegogo.com/1.1/campaigns/YOUR_ID/contributions.json?api_token=API_TOKEN&per_page=200")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Indiegogo

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	// This should show the Donor
	fmt.Fprintf(w, "Donors: %s", record.Response)

}
