package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("0x%x\n", hdr.Data)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dummpBytesArray((*p)[:])

	// 下标操作
	var s1 = "中国人"
	fmt.Printf("0x%x\n", s1[0])

	// 字符迭代
	for i := 0; i < len(s1); i++ {
		fmt.Printf("index:%d, value:0x%x\n", i, s1[i])
	}

	var arr1 [5]int
	fmt.Println("len:", len(arr1))
	fmt.Println("size:", unsafe.Sizeof(arr1))

	var arr2 = [...]int{
		99: 99,
	}
	fmt.Println("len:", len(arr2))

	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := arr3[3:7:9]
	sl[0] += 10
	fmt.Println("arr3[3]", arr3[3])
}

func dummpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c", b)
	}
	fmt.Printf("]\n")
}
