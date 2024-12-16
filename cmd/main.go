package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	httpRequester "github.com/andytandy24/internal"
)

func main() {
	var wg sync.WaitGroup

	// CLI flags
	url := flag.String("u", "http://127.0.0.1", "Specifies the target url")

	method := flag.String("X", "GET", "Specifies the http method that should be used")

	numRequests := flag.Int("n", 20, "Defines the amount of simultaneous requests sent to a target")

	data := flag.String("d", "", "Specifies the data that is meant to be sent")

	debug := flag.Bool("D", false, "Toggles debug mode, is false by default")

	flag.Parse()

	// Requests URL
	for i := 0; i < *numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			httpRequester.MakeRequest(*method, *url, *data)

			if *debug {
				fmt.Println("Completed at:", time.Now().String())
			}
		}()
	}
	wg.Wait()
}
