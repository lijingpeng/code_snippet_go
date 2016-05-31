package main

import (
	"fmt"
)

type LogicExpression struct {
	// 与查询， 必须都存在
	MustLabels []string

	// 或查询， 有一个存在即可
	ShouldLabels []string

	// 非查询， 不包含
	NotInLabels []string
}

func test_interface(i int, a ...LogicExpression) {
	fmt.Println(i)
	if a != nil {
		fmt.Println(a[0].MustLabels)
	}
}

func main() {
	test_interface(1)
	test_interface(2, LogicExpression{})

	m1 := []string{"1", "2"}
	s1 := LogicExpression{MustLabels: m1}
	test_interface(3, s1)
}
