package main 
import "fmt"

func main() {
    msg := make(chan string, 2)	// ->have number for Buffered Channel
 
    msg <- "buffered"
    msg <- "channel"
 
    a := <- msg
    fmt.Println(a)

    a = <- msg
    fmt.Println(a)
}