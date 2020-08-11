package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/* 1 */
	var t, f bool = true, false
	fmt.Println(t, f)

	fmt.Printf("%T and %v and %v bytes\n", t, t, unsafe.Sizeof(f))
}
