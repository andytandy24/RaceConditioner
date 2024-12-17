package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	internal "github.com/andytandy24/internal"
)

func main() {
	var wg sync.WaitGroup

	// CLI flags
	url := flag.String("u", "http://127.0.0.1", "Specifies the target url")
	method := flag.String("X", "GET", "Specifies the http method that should be used")
	numRequests := flag.Int("n", 20, "Defines the amount of simultaneous requests sent to a target")
	data := flag.String("data", "", "Specifies the data that is meant to be sent")
	headers := flag.String("H", "", "Specifies the headers of the request")
	debug := flag.Bool("D", false, "Toggles debug mode, is false by default")

	flag.Parse()

	// Requests URLs
	wg.Add(*numRequests)
	for i := 0; i < *numRequests; i++ {
		go func() {
			defer wg.Done()
			internal.MakeRequest(*method, *url, *data, internal.ParseHeaders(*headers))

			if *debug {
				fmt.Println("Completed at:", time.Now().String())
			}
		}()
	}
	wg.Wait()
}
