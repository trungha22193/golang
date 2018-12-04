package main
import "fmt"

func main() {
	fmt.Println(recursive(5,0))
}

func recursive(x int, total int) int {
	if (x == 0) {
		return total
	}
	total += x
	x--
	return recursive(x, total)
}
