package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "42" || arg == "piscine" || arg == "piscine 42" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
