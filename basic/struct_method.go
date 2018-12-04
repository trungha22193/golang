package main 
import ("fmt")

type Circle struct {
	x,y,r float64
	
}

func (c *Circle) area() float64 {	// → assign (c *Circle) to function of Method in Struct
	return c.y * c.r
}

func main() {
	c := Circle{1,2,3}
	fmt.Println(c.y)
	fmt.Println(c.area())
}
