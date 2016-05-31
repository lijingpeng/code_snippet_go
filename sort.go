package main

// import (
// 	"fmt"
// 	"github.com/AlasdairF/Sort/Uint32"
// 	// "sort"
// )

// func main() {
// 	a := []uint32{1, 2, 45, 4, 35, 9, 8}

// 	fmt.Println(sortUint32.Desc(a))
// }
import (
	"fmt"
	"github.com/AlasdairF/Sort/Uint64"
)

func main() {
	list := []uint64{10, 44, 1, 7, 4, 0, 9, 0, 3, 65, 38}
	sortUint64.StableDesc(list)
	fmt.Println(list)
}
