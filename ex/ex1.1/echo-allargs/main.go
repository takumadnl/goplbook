// 練習問題1.1
// echo-allargsはコマンドライン引数とプログラムを起動したコマンド名を表示する
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
