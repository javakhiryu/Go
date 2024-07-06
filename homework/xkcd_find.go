package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	URL        string = "https://xkcd.com/"
	targetFile string = "/info.0.json"
	inputFile  string = "xkcd_comics.json"
)

// Response represents the JSON structure from the XKCD API
type Response struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading JSON to file:", err)
		os.Exit(-1)
	}

	defer jsonFile.Close()

	// Read the JSON file
	var responses []Response
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&responses)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding JSON:", err)
		os.Exit(-1)
	}

	args := os.Args[1:]
	joinedArgs := strings.Join(args, " ")

	// Print the data for demonstration
	for _, response := range responses {
		if strings.Contains(response.Title, joinedArgs) || strings.Contains(response.Transcript, joinedArgs) {
			fmt.Printf("URL: %s%d%s\nDate: %s-%s-%s\nTitle: %s", URL, response.Num, targetFile, response.Day, response.Month, response.Year, response.Title)
			break
		}
	}
}
