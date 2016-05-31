package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)

	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a[0], a[1])

}
