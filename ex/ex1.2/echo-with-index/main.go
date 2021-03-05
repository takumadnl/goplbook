// 練習問題1.2
// echo-with-indexはコマンドライン引数を表示する
// 個々の引数のインデックスと値の組を１行ごとに表示する
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println("index:", i, ", argument:", arg)
	}
}
