package main
import "fmt"

func main() {
	for i:=1;i<=10;i++ {
		switch i {
			case 1: fmt.Println("Mot")
			case 2: fmt.Println("Hai")
			case 3: fmt.Println("Ba")
			default: fmt.Println("Khong")
		}
	}
}
