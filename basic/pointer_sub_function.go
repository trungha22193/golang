package main 
import "fmt"

func main() {
	x := 1
	pointer(&x)
	fmt.Println(x)
}

func pointer(xPnt *int) {
	*xPnt = 2
}
