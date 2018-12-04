package main
import "fmt"

func main() {
	defer func1()
	func2()
}

func func1() {
	fmt.Println("Func1")
}

func func2() {
	fmt.Println("Func2")
}
