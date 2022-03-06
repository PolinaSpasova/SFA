package main

import (
	"fmt"
	"math/rand"
	"time"
)

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	m:=make(map[string][]int)
	for idx,name :=range keySlice{
	m[name]= append(m[name],valueSlice[idx])
		}
	return m
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100
	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}
	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	return cities, prices
}

func main() {
	keySlice,valueSlice:=citiesAndPrices()
	m:=groupSlices(keySlice,valueSlice)
	for k,v := range m{
		fmt.Println(k, v)
	}
}