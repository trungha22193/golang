package main
import "fmt"

func main() {
	slide := []float64{1,2,3,4,5}
	fmt.Println(average(slide))
}

func average(input []float64) float64 {	// → should return exact value
	total := 0.0				// → assign value to variable right away
	for _, value := range input {
		total += value
	}
	return total / float64(len(input))
}
