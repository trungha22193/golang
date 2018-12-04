package main 
import "fmt"

type Fraction struct {
	x, y int
}

func (e *Fraction) Error() string {
	return fmt.Sprintf("Can't divide %d by %d", e.x, e.y)		// → Sprintf, not Printf
}

func divide(x, y int) (int, error) {		// return “error” not “s”
	if y == 0 {
		err := Fraction{x, y}
		return -1, &err			// → use &err because *Fraction
	} 
	return x/y, nil
}

func main() {
	x := 5
	y := 0
	result, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
