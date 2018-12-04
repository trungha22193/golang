package main 
import "fmt"

func main() {
	x := 0
	increase := func() int {
		x++
		return x
	}
	fmt.Println(increase())
	fmt.Println(increase())
}
