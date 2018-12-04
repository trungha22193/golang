package main
import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	for i:=0;i<10;i++ {
		m["key"] = i + 1
	}
	fmt.Println(m)
	delete(m,"key")
	fmt.Println(m)
}
