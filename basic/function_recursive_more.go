package main 
import "fmt"

func recursive(arg []int) int {
	if len(arg) == 1 {
		return arg[0]
	} 
	return arg[0] + recursive(arg[1:])
}

func main() {
	arr := []int{1,2,3,4,5}
	fmt.Println(recursive(arr))
}
