package internal

import (
	"log"
	"net/http"
	"strings"
)

func MakeRequest(method string, url string, data string, headers map[string]string) {
	var req *http.Request
	var err error

	if data != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Fatalf("Failed to establish HTTP request: %s", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("The response returned: %s", resp.Status)
	}
}
