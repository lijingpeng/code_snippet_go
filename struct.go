package main

import (
	"fmt"
)

// func deepCopy(dst, src interface{}) error {
// 	var buf bytes.Buffer
// 	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
// 		return err
// 	}
// 	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
// }

type key struct {
	docs []uint32
	fre  []float32
}

func struct_nil(i int, k ...key) {
	fmt.Println(len(k))
	if len(k) > 0 {
		fmt.Println(k[0].docs)
	}
}

func main() {
	var k1 key
	array := []uint32{1, 2, 3}
	k1.docs = array
	// fmt.Println(k1.docs[1])
	struct_nil(1, k1)

	k2 := key{}
	// k2.docs = array
	// fmt.Println(k2.docs[1])
	struct_nil(2, k2)

}
