package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	j int64
}

func main() {
	n := Num{i: "abc", j: 4}
	fmt.Printf("struct: %p  ->  Num.i: %p \n", &n, &n.i)
	ptr := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + uintptr(0)))
	fmt.Printf("%v \n", unsafe.Offsetof(n.j))
	fmt.Printf("%v    %v", *ptr, n.j)
}

