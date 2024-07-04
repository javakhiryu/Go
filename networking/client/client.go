package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL string = "https://jsonplaceholder.typicode.com"

type todo struct {
	ID        int    `json"id`
	Title     string `json:"title"`
	Completed bool   `bool:"completed"`
}

func main() {
	resp, err := http.Get(URL + "/todos/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading response body:", err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading response body:", err)
			os.Exit(-1)
		}
		fmt.Printf("%#v\n", item)
	} else {
		fmt.Printf("Received non-OK HTTP status: %s\n", resp.Status)
	}
}
