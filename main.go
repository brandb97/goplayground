package main

import (
	"os"
)

func main() {
	switch {
	case os.Args[0] == "client":
		client()
	case os.Args[0] == "server":
		server()
	}
}
