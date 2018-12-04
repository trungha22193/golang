package main
import "fmt"

func main() {
	const x string = "hello world"
    	fmt.Println(x)

    	x = "test"   			// → can’t assign to CONST variables
}
