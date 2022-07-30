package main

import "testing"
import "fmt"
import "net"

func TestClosure(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}

func TestIp(t *testing.T) {
	ips, error := net.LookupIP("fedora-yld")
	fmt.Printf("%v %v\n", ips, error)
}
