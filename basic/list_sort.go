package main 
import "fmt"
import "sort"

func main() {
    x := []int{4,6,7,1,2}
    y := []float64{3.5, 7.6, 8.93, 5.23, 7.609}

    fmt.Println(x)
    sort.IntSlice.Sort(x)
    fmt.Println(x)

    fmt.Println(y)
    sort.Float64Slice.Sort(y)
    fmt.Println(y)
}
