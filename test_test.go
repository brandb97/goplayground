package main

import "testing"
import "fmt"
import "net"
import "os"

func TestIp(t *testing.T) {
	hn, _ := os.Hostname()
	ips, _ := net.LookupIP(hn)
	fmt.Printf("%v\n", ips)
}
