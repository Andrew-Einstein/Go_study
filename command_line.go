package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "KIKI"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("你好", who)
}
