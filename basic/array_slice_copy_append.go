package main
import "fmt"

func main() {
	slide1 := []int{2,3,4}
	slide2 := make([]int,2)
	copy(slide2, slide1)
	fmt.Println(slide1)
	fmt.Println(slide2)

	slide2 = append(slide2, 5, 6)
	fmt.Println(slide2)
}
