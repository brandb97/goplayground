package main

import (
	"os"
)

const (
	serverIp   = "127.0.0.1"
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
