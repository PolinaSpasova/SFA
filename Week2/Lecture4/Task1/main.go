package main

import "fmt"

func daysInMonth(month, year int) (int,bool){
	var days int 
	ok:=true
	if year>0 && year<5000{
		switch month {
		case 1,3,5,7,8,10,12:
			days=31
		case 4,6,9,11:
			days=30
		case 2:
			if year%4==0 {
				days=29
			}else{
				days=28
			}
		default: 
			ok=false
		}
	}else{
		ok=false
	}
	return days,ok;
}

func main(){
	if days,err:=daysInMonth(2,1985); err{
		fmt.Printf("This month has %d days.",days)
	}else{
		fmt.Printf("The month or the year you choose is out of range!")
	}
}