package slice

import (
	"fmt"
	"unsafe"
)

func PrintDynamicGrow() {
	// slice unbind orig array
	// slice is "inference"
	u := [...]int{11, 12, 13, 14, 15}
	fmt.Printf("init array:%v, address:%p\n", u, &u)
	s := u[1:3]
	s = append(s, 24)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	fmt.Printf("after append 24, array:%v, address:%p\n", u, &u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s = append(s, 25)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	fmt.Printf("after append 25, array:%v, address:%p\n", u, &u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s[0] = 0
	fmt.Printf("after reassign first element of slice, array:%v, address:%p\n", u, &u)
	fmt.Printf("after reassign first element of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	fmt.Println()

	s2 := make([]int, 0, 5)
	s2 = append(s2, []int{2, 3, 4, 5, 6}...)
	underlyingA := unsafe.Pointer(&s[0])
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s2), cap(s2), s2)
	fmt.Printf("underlying array address: %v\n", underlyingA)

	s2 = append(s2, 7)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s2), cap(s2), s2)
	fmt.Printf("after append 7, underlying array address:%v\n", underlyingA)
	fmt.Printf("after append 7, slice(len=%d, cap=%d): %v\n", len(s2), cap(s2), s2)

	s2 = append(s2, 8)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s2), cap(s2), s2)
	fmt.Printf("after append 8, underlying array address:%v\n", underlyingA)
	fmt.Printf("after append 8, slice(len=%d, cap=%d): %v\n", len(s2), cap(s2), s2)
}
