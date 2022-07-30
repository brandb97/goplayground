package main

import "testing"
import "fmt"
import "net"
import "os"

func TestClosure(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}

func TestIp(t *testing.T) {
	hn, _ := os.Hostname()
	ips, _ := net.LookupIP(hn)
	fmt.Printf("%v\n", ips)
}
