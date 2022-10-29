package main

import (
	"fmt"
)

func tryDefer() {
	// 順序會是 30, 29, ..., 0
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func tryDefer2() {
	a := 1
	defer func(b int) {
		fmt.Println("defer b", b)
	}(a + 1)
	a = 99
}

var g = 100

// keep return value => do defer => return
func f() (r int) {
	defer func() {
		g = 200
	}()
	fmt.Printf("f: g = %d\n", g)
	return g
}

func main() {
	tryDefer2()
	fmt.Println("============================================")
	fmt.Println(f(), g)
	fmt.Println("============================================")
	tryDefer()

}
