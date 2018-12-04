package main
import "fmt"

func main() {
	m := make(map[string]int)
	m["key"] = 10

	value, flag := m["key"]
	fmt.Println(value, flag)

	value, flag = m["key1"]
	fmt.Println(value, flag)
}
