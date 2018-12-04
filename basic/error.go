package main 
import "fmt"
import "errors"					// → has “s”

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return -1, errors.New("Can't divide by 0")
	} else {
		return x/y, nil
	}
}

func main() {
	x := 5.0
	y := 0.0
	result, err := divide(x, y)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Result =", result)
	}
}
