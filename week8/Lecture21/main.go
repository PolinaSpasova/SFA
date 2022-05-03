package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type TopStories struct {
	Title string `json:"title,omitempty"`
	Score int    `json:"score,omitempty"`
}

type TopJSONRes struct {
	Top_Stories [10]TopStories
}

func GetTop() []int {
	u := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatal(err)
	}
	httpRes, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer httpRes.Body.Close()

	var res []int
	json.NewDecoder(httpRes.Body).Decode(&res)
	return res[:10]
}

func HandleTop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := GetTop()
		var result [10]TopStories
		var final TopJSONRes
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				u := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", str[i])
				req, err := http.NewRequest("GET", u, nil)
				if err != nil {
					log.Fatal(err)
				}
				httpRes, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				defer httpRes.Body.Close()
				json.NewDecoder(httpRes.Body).Decode(&result[i])
				log.Println(result[i].Title)
				wg.Done()
			}(i)
		}
		wg.Wait()
		final.Top_Stories = result
		err := json.NewEncoder(w).Encode(&final)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/top", HandleTop())
	http.ListenAndServe(":9000", mux)
}
