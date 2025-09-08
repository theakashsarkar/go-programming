package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	c := a

	fmt.Println("Contents: ")
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(c, a))

	fmt.Println("\nMemory addresses of first elements:")
	fmt.Printf("a: %p\n", &a[0])
	fmt.Printf("b: %p\n", &b[0])
	fmt.Printf("c: %p\n", &c[0])

	fmt.Println("\nMemory comparison:")
	fmt.Println("a vs b:", &a[0] == &b[0])
	fmt.Println("a vs c:", &a[0] == &c[0])
}
