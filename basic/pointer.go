package main 
import "fmt"

func main() {
	var x int = 1
	var y *int		// → define pointer
	y = &x			// → assign address of variable x to pointer y
	fmt.Println(x)
	fmt.Println(y)		// → print address of variable x
	x = 2			
	fmt.Println(*y)		// → print value of Pointer y
}
