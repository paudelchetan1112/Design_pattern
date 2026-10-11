package main

import (
	"fmt"
	"math"
)
type Shape interface{
	area() float64
	perimeter() float64


}
type Rectangle struct{
	length float64
	breadth float64
}
func(r Rectangle)area()float64{
	return r.length*r.breadth
}
func(r Rectangle)perimeter()float64{
	return 2*(r.length+r.breadth)
}

type Circle struct{
	radius float64
}
func(c Circle)area()float64{
	return math.Pi*c.radius*c.radius
}
func (c Circle)perimeter()float64{
	return 2*math.Pi*c.radius
}
func Area(s Shape)float64{
	return s.area()
}
func Perimeter(s Shape)float64{
	return s.perimeter()
}
func main(){
	r1:=Rectangle{length: 10, breadth: 20}
	c1:=Circle{radius: 5}
	fmt.Println("Area of Rectangle R1:", Area(r1))
	fmt.Println("Area of circle:", Area(c1))

}