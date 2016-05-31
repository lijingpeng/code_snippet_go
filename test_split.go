package main

import (
	"fmt"
	"strings"
	// "log"
)

func main() {
	s := "268996,,广州千思电脑,http://shop268996.taobao.com/,"
	a := strings.Split(s, ",")
	fmt.Println("length:", len(a))
	for _, i := range a {
		fmt.Println(i)
	}
}
