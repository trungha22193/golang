package main
import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()				// → should have “()” for closure with “defer”
	panic("Have error")
	fmt.Println("Revived")		// → not execute because panic() function
}
