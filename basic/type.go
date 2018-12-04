package main
import "fmt"

func main(){
	fmt.Println("Hello World !")

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) 		// → return ASCII code, not real number
	fmt.Println("Hello" + " " + "World")
}
