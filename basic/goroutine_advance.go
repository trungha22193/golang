package main
import "fmt"

func f(n int) {
	for i:=0;i<5;i++ {
		fmt.Println(n,":",i)
	}
}

func main() {
	for i:=0;i<5;i++ {
		go f(0)
	}
	var input string
	fmt.Scanln(&input)
}
