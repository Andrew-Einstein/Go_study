package main

import (
	"fmt"
	"os"
	"strings"
)

// 命令行参数
func main() {
	who := "World!"
	if len(os.Args) > 1 {
		// os.Args[0] is CommandLine.exe or CommandLine
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello", who)
}
