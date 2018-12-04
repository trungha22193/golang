package main
 
import "fmt"
 
func main() {
    msg := make(chan string, 2)
 
    msg <- "buffered"
    msg <- "channel"
 
    go func() {
    	for {
    		a := <- msg
    		fmt.Println(a)
    	}
    }()

    var input string
    fmt.Scanln(&input)
}
