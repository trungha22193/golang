package main 
import "fmt"

type Fraction struct {
    x, y float64
}

func (f *Fraction) Error() string {
    return fmt.Sprintf("Can't divide %f by %f", f.x, f.y)
}

func divide(x, y float64) (float64,error) {
    if y == 0 {
        err := Fraction{x,y}
        return -1, &err
    }
    return x/y, nil
}

func main() {
    x := 5.0
    y := 0.0
    result, err := divide(x,y)
    if err != nil {
        frac, flag := err.(*Fraction)

        fmt.Println(err)
        fmt.Println("Flag: ", flag)
        if flag {
            fmt.Println("Frac.x & Frac.y :", frac.x, frac.y)    
        }
    } else {
        fmt.Println(result)
    }
}
