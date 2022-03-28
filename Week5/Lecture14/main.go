package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func main() {
	var concurrencyLevel int
	flag.IntVar(&concurrencyLevel, "c", 2, "maximum number of concurrent connections")
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) < 1 {
		fmt.Println("multiping [-c] <filePaths>")
		fmt.Println("There are no filePaths")
		flag.PrintDefaults()
		return
	}

	processQueue := make(chan string, concurrencyLevel)
	done := make(chan string)

	go func() {
		for _, urlToProcess := range filePaths {
			processQueue <- urlToProcess
			go func(url string) {
				if err := pingURL(url); err != nil {
					fmt.Println("There is an error: ", err)
					return
				}
				<-processQueue
				done <- url
			}(urlToProcess)
		}
	}()

	for range filePaths {
		log.Println("Done:", <-done)
	}

}

// func fetch(urls []string, concurrency int) {
// 	processQueue := make(chan string, concurrency)
// 	done := make(chan string)
// 	go func() {
// 		for _, urlToProcess := range urls {
// 			processQueue <- urlToProcess
// 			go func(url string) {
// 				//	pingURL(url)
// 				<-processQueue
// 				done <- url
// 			}(urlToProcess)
// 		}
// 	}()
// 	for range urls {
// 		log.Println("Done:", <-done)
// 	}
// }
