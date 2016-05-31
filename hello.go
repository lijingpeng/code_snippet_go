package main

import (
	"fmt"
)

func test_args(i int, args ...string) {
	for _, v := range args {
		fmt.Println(i, v)
	}
}

func test_slice(a []int) {
	fmt.Println(a[0])
}

func bool_func() bool {
	return false
}

func main() {
	fmt.Println("hello.go")
	t := []int{1, 2, 4, 9, 10}
	idx := 3
	fmt.Println(t[idx:])
	fmt.Println(t[idx+1:])
	t = append(t, 0)
	copy(t[idx+1:], t[idx:])
	t[idx] = 5
	fmt.Println(t)

	fmt.Println("-----------")
	token := []string{"中国", "城市"}
	label := []string{"香港", "迪斯尼"}
	keywords := make([]string, len(token)+len(label))
	copy(keywords, token)
	copy(keywords[len(token):], label)
	for i := range keywords {
		fmt.Println(keywords[i])
	}
	fmt.Println("keywords length:", len(keywords))
	fmt.Println("-----------")
	b := false
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Println("i:j->", i, j)
			if i == 5 {
				b = true
				break
			}
		}
		if b {
			fmt.Println("bbbbb")
		}
	}
	fmt.Println("-----------")
	test_args(1)
	test_args(2, "a")
	test_args(3, "b", "c")

	fmt.Println("-----------")
	test_slice(t[0:])
	test_slice(t[1:])

	mm := make([]int, 0)
	fmt.Println(len(mm))

	bol := true
	bol = bool_func()
	fmt.Println(bol)
}
