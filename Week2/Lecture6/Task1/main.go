package main

import "fmt"

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func add(l *MagicList, value int) {
	new := Item{Value: value}
	if l.LastItem == nil {
		l.LastItem = &new
	} else {
		new.PrevItem = l.LastItem
		l.LastItem = &new
	}
}

func toSlice(l *MagicList) []int {
	var res []int
	for l.LastItem != nil {
			res = append([]int{l.LastItem.Value},res...)
		//	res=append(res, l.LastItem.Value)
			l.LastItem=l.LastItem.PrevItem
		}
	return res
}


func main() {
	l := &MagicList{}
	add(l, 10)
	add(l, 22)
	add(l, 44)
	add(l, 21)
	add(l, 99)
	add(l, 11)
	add(l, 18)
	add(l, 2)
	add(l, 5)
	add(l, 17)

	fmt.Println(toSlice(l))
}
