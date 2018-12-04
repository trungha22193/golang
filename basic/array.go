package main
import "fmt"

func main() {
	var arr [5]float64
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500

	var total float64 = 0
	for i:=0;i<len(arr);i++ {
		total += arr[i]
	}

	fmt.Println(total / float64(len(arr)))
}
