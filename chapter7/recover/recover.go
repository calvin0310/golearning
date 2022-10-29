package main

import (
	"fmt"
)

func tryRecover() {
	// defer func() {}() 相當於 defer errHandler()
	defer func() {
		// recover return interface，所以不一定是error
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			//panic(r)

			// (3)
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	// (1)
	//panic(errors.New("this is an error"))

	// (2)
	b := 0
	a := 5 / b
	fmt.Println(a)

	// (3)
	//panic(123)
}

func main() {
	tryRecover()
}
