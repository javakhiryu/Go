package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	URL        string = "https://xkcd.com/"
	targetFile string = "/info.0.json"
	outputFile string = "xkcd_comics.json"
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

func fetchComic(Pagecounter int, wg *sync.WaitGroup, responses *[]Response, mu *sync.Mutex, NFCounter *int32) {
	defer wg.Done()

	resp, err := http.Get(URL + strconv.Itoa(Pagecounter) + targetFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		atomic.AddInt32(NFCounter, 1)
		return
	}

	atomic.StoreInt32(NFCounter, 0)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading response body:", err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unmarshalling JSON:", err)
		return
	}

	mu.Lock()
	*responses = append(*responses, response)
	mu.Unlock()
}

func main() {
	var responses []Response
	var NFCounter int32
	var Pagecounter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for {
		wg.Add(1)
		go fetchComic(Pagecounter, &wg, &responses, &mu, &NFCounter)
		fmt.Println(Pagecounter)
		Pagecounter++

		wg.Wait()
		if atomic.LoadInt32(&NFCounter) >= 2 {
			break
		}
	}

	// Marshal the slice of responses into JSON
	jsonData, err := json.MarshalIndent(responses, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshalling to JSON:", err)
		os.Exit(-1)
	}

	// Write the JSON data to a file
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing JSON to file:", err)
		os.Exit(-1)
	}

	fmt.Println("Data successfully saved to", outputFile)
}
