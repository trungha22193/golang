package main 
import "fmt"

type Circle struct {
	x,y,r uint32
}

func main() {
	c := new(Circle)
	c.x = 1
	c.y = 2 
	fmt.Println(c.y)
}
