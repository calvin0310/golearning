package main

import (
	"bufio"
	"fmt"
	"golearning/chapter7/defer/fib"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 怕忘記所以就近寫Close
	defer file.Close()

	// Buffer比較快，先寫到RAM，等累積到一定的量再一次寫入HDD
	writer := bufio.NewWriter(file)
	// 將writer導入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
