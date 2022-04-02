package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) ([]string, error) {
	var res []time.Time
	var new []string
	for i:=0;i<len(dates);i++{
		parsed,err:=time.Parse(format,dates[i])
		if err!=nil{
			return nil,err
		}
		res=append(res,parsed)
	}
	sort.Slice(res,func(i, j int) bool { return res[i].Before(res[j])})
	for j:=0;j<len(res);j++{
		new=append(new, res[j].Format("Jan-02-2006"))
	}
	return new,nil
}

func main() {
	dates := []string{"Dec-03-2021","Sep-14-2008","Mar-18-2022"}
	format:= "Jan-02-2006"
	fmt.Println(dates)
	res, err:=sortDates(format,dates...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}