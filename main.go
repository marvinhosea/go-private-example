package main

import (
	"fmt"
	"github.com/marvinhosea/go-filter"
)

func main() {
	fmt.Println("Go private module Example")

	intArray1 := []int{1, 2, 3, 4, 5, 5}
	intArray2 := []int{0, 9, 45, 5}
	result := filter.Filter(intArray1, intArray2)
	fmt.Println(result)
}
