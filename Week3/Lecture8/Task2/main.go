package main

import("fmt")

type Square struct{
	a int
}

func NewSquare(a int) *Square{
	return &Square{a:a}
}

func (r *Square) Area() float64 {
	return float64(r.a)*float64(r.a)
}

type Circle struct{
	r int
}

func NewCircle(r int) *Circle{
	return &Circle{r:r}
}

func (r *Circle) Area() float64 {
	return float64(r.r)*float64(r.r)*3.14
}

type Shape interface{
	Area() float64
}

type Shapes []Shape

func (s Shapes) LargestArea() float64{
	max:=s[0].Area()
	for i:=1;i<len(s);i++{
		if max<s[i].Area() {
			max=s[i].Area()
		}
	}
	return max
}

func main() {
	var try Shapes
	try=append(try,NewCircle(3),NewSquare(6),NewSquare(7),NewCircle(12))
	fmt.Println(try.LargestArea())
} 