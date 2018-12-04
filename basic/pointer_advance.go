package main 
import "fmt"

func main() {
	x := 1
	y := new(int)		// → define new Pointer
	y = &x
	pointer(y)
	fmt.Println(*y)
}

func pointer(xPnt *int) {
	*xPnt = 2
}
