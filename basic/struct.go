package main 
import "fmt"

type Circle struct {
    x, y, r float64 		// → define Attribute without “var”
}

func main() {
	var c Circle		
	c.x = 1
	c.y = 2
	c.r = 3
	fmt.Println(c.x)

	d := Circle{4,5,6}
	fmt.Println(d.y)
}
