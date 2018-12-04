package main
import "fmt"

func main() {
	x := [5]float64{100,200,300,400,500}

	var total float64 = 0

	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}
