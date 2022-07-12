package main

import "fmt"
import "os"

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
