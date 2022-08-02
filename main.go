package main

import (
	"os"
	"strings"
)

const (
	serverIp   = "47.107.108.217"
	serverPort = ":1234"
	MAXINPUT   = 1024
)

func main() {
	switch {
	case strings.HasSuffix(os.Args[0], "client"):
		client()
	case strings.HasSuffix(os.Args[0], "server"):
		server()
	}
}
