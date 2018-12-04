package main 
import "fmt"
import "sort"

type Person struct {
     name string
     age int
}

type ByAge []Person

func (this ByAge) Len() int {
    return len(this)
}

func (this ByAge) Less(i, j int) bool {
    return this[i].age < this[j].age
}

func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i] 
}

func main() {
    kids := []Person{
        {"A", 12},
        {"B", 10},
    }
    sort.Sort(ByAge(kids))
    fmt.Println(kids)
}