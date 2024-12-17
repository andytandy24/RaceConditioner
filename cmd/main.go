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
	numRequests := flag.Int("n", 20, "Defines the amount of simultaneous requests sent to a target")
	debug := flag.Bool("D", false, "Toggles debug mode, is false by default")
	cookie := flag.String("cookie", "", "Specifies the cookie of the request")
	data := flag.String("data", "", "Specifies the data that is meant to be sent")
	headers := flag.String("H", "", "Specifies the headers of the request")
	method := flag.String("X", "GET", "Specifies the http method that should be used")
	url := flag.String("u", "http://127.0.0.1", "Specifies the target url")

	flag.Parse()
	headerMap := internal.ParseHeaders(*headers)
	cookieMap := internal.ParseHeaders(*cookie)

	req := internal.MakeRequest(*method, *url, *data, headerMap, cookieMap)
	wg.Add(*numRequests)

	for i := 0; i < *numRequests; i++ {
		go func() {
			defer wg.Done()
			internal.SendRequest(req)

			if *debug {
				fmt.Println("Completed at:", time.Now().String())
			}
		}()
	}

	wg.Wait()
}
