package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{{
		name:     "index",
		path:     "/api",
		expected: "Hello world!",
	},
		{
			name:     "healthCheck",
			path:     "/healthCheck",
			expected: "ok",
		},
	}
	mux := http.NewServeMux()
	setupHandler(mux)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tc.path)
			respBody, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			if string(respBody) != tc.expected {
				t.Errorf("Expected %q, got %q",
					tc.expected, string(respBody))
			}

		})
	}
}
