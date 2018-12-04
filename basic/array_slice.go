package main
import "fmt"

func main() {
	x := [5]float64{1,2,3,4,5}
	slide := make([]float64, 5)	// → create slide with 5 value 0

	fmt.Println(x)
	fmt.Println(slide)

	slide = x[0:5]
	fmt.Println(slide[:])		// → return all 
	fmt.Println(slide[1:3])		// → return number of position 1 to <3
	fmt.Println(slide[1:])		// → return value from position 1 to end
	fmt.Println(slide[:4])		// → return value from start to <4
}
