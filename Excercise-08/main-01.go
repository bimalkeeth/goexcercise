package main

import (
	"fmt"
	"unsafe"
)

func ReadMemory(ptr unsafe.Pointer, size uintptr) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))))
	}
	return out
}

type goString struct {
	elements []byte
	len      int
}

func main() {
	s := []byte("Hello World")
	var stringExample = "Hello World"
	var anotherStringExample = "Hello World"
	var goString = goString{s, 12}
	sz := unsafe.Sizeof(stringExample)
	fmt.Println(sz)

	fmt.Println(unsafe.Pointer(&stringExample))
	fmt.Println(unsafe.Pointer(&anotherStringExample))

	stringExample = anotherStringExample

	fmt.Println(unsafe.Pointer(&stringExample))
	fmt.Println(unsafe.Pointer(&anotherStringExample))

	n := unsafe.Pointer(&goString.elements[0])
	fmt.Println(ReadMemory(n, 11))
	fmt.Println(string(ReadMemory(n, 11)))

}
