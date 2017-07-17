package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Indiegogo struct for donors names from response
type Indiegogo struct {
	Response []struct {
		By string `json:"by"`
	} `json:"response"`
}

var record Indiegogo

func main() {
	updateNames()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// handler echoes the donors
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Unsupported HTTP Method", http.StatusMethodNotAllowed)
		return
	}

	donors, err := json.Marshal(record.Response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(donors)

	return

}

func updateNames() {
	url := "https://api.indiegogo.com/1.1/campaigns/YOUR_ID/contributions.json?api_token=API_TOKEN&per_page=200"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
}
