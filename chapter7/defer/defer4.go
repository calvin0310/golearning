package main

import (
	"bufio"
	"fmt"
	"golearning/chapter7/defer/fib"
	"os"
)

func writeFile(filename string) {
	// O_CREATE：檔案不存在就建立檔案, O_EXCL：檔案存在就回傳失敗
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("this is a custom error")
	// 因為這個自定義的不是pathError, 所以會直接panic

	// this error is a pathError
	// 看IDE說明
	// 這種方式的好處就是將已知類型的錯誤處理掉，未知類型的錯誤印出來
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
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
