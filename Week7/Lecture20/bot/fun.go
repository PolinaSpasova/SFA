package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Start() {
	for {
		f := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
		var drink string
		fmt.Println("What would you like to drink?")
		fmt.Scanln(&drink)
		if drink == "nothing" {
			break
		}
		newOrder := f + drink
		req, err := http.NewRequest("GET", newOrder, nil)
		if err != nil {
			log.Fatal(err)
		}
		httpRes, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer httpRes.Body.Close()

		var res BartenderBot
		json.NewDecoder(httpRes.Body).Decode(&res)
		if res.Drinks==nil{
			fmt.Println("We don't have that drink!")
			continue
		}
		fmt.Println(res.Drinks[0].StrDrink)
		recipe := strings.Split(res.Drinks[0].StrInstructions, ".")
		 for _, v := range recipe {
		 	fmt.Print(v)
		}
		println()
	}
}