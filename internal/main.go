package internal

import (
	"log"
	"net/http"
)

func MakeRequest(method string, url string, data string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("Failed to create http request: %s", err)
	}

	if data != "" {
		//TBD
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send http request: %s", err)
	}
}
