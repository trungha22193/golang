package main
import "fmt"

func sum(args ...int) int {		// → define represent variables for multi-values
	total := 0
	for _, value:= range args {
		total += value
	}
	return total 
}

func main() {
	fmt.Println(sum(1,2,3))
}
