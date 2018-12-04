package main
import "fmt"

type Shape interface {
	area() float64		
}

type Circle struct {
	x,y,r float64
}

func (c *Circle) area() float64 {		// → (**) remove * on both
	return c.x * c.y
}

type Rectangle struct {
	l,w float64
}

func (r *Rectangle) area() float64 {		// → (**) remove * on both
	return r.l * r.w
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

func main() {
	c := Circle{1,2,3}
	r := Rectangle{4,5}

	fmt.Println(c.area())
	fmt.Println(r.area())

	fmt.Println(totalArea(&c,&r))		// → can use totalArea(c,r) without * at (**) above
}
