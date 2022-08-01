package main

import (
	"os"
)

const (
	serverIp   = "47.107.108.217"
	serverPort = ":1234"
	MAXINPUT   = 1024
)

func main() {
	switch {
	case os.Args[0] == "./client":
		client()
	case os.Args[0] == "./server":
		server()
	}
}
