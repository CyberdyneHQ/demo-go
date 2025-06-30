package main

import (
	"fmt"
	"unsafe"
)

type Fake struct{}

func (Fake) Good() {}

func unsafeCode() {
	unsafeM := Fake{}
	unsafeM.Good()
	intArray := [...]int{1, 2}
	fmt.Printf("\nintArray: %v\n", intArray)
	intPtr := &intArray[0]
	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n", intPtr, *intPtr)
	addressHolder := uintptr(unsafe.Pointer(intPtr)) + unsafe.Sizeof(intArray[0])
	intPtr = (*int)(unsafe.Pointer(addressHolder))
	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n\n", intPtr, *intPtr)
	// Create a context with a value
	ctx := context.WithValue(context.Background(), "somekey", "somevalue")

	// Call the function with the context
	result := GetSomethingWithContext("example", ctx)

	// Print the result
	log.Println(result)
}
