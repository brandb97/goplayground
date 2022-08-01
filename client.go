package main

import (
	"bufio"
	"fmt"
	"goplay/message"
	"log"
	"net"
	"os"
)

func client() {
	conn, err := net.Dial("tcp", serverIp+serverPort)
	if err != nil {
		log.Fatalln(err)
	}
	name := MakeName()
	chat(name, conn)
}

func MakeName() string {
	hn, _ := os.Hostname()
	ipSlice, _ := net.LookupIP(hn)

	// find ipv4 addr in ipSlice
	ipv4 := func() string {
		for i := range ipSlice {
			for j := range ipSlice[i] {
				if ipSlice[i][j] == byte('.') {
					return ipSlice[i].String()
				}
			}
		}
		return ""
	}()
	pid := os.Getpid()
	return ipv4 + fmt.Sprint(pid)
}

func chat(name string, conn net.Conn) {
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Your name(default %s): ", name)
	input.Scan()
	if buf := input.Text(); buf != "" {
		name = buf
	}
	fmt.Printf("login %s%s using name: %s\n", serverIp, serverPort, name)

	fmt.Printf("Peer's name: ")
	input.Scan()
	peer := input.Text()
	fmt.Printf("Chat with peer: %s\n", peer)

	var connCh chan []byte = make(chan []byte)
	var inputCh chan []byte = make(chan []byte)
	go func() {
		for {
			var data []byte = make([]byte, MAXINPUT)
			n, err := conn.Read(data)
			if err != nil {
				close(connCh)
				log.Fatal(err)
			} else {
				connCh <- data[:n]
			}
		}
	}()
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			inputCh <- []byte(input.Text())
		}
	}()

	sendMsg := message.SendMsg{
		Name:       name,
		TargetName: peer,
	}
	for {
		var receiveMsg message.ReceiveMsg
		select {
		case sendMsg.Body = <-inputCh:
			data := sendMsg.Encode()
			_, err := conn.Write(data)
			if err != nil {
				log.Fatalln(err)
			}
		case receiveData := <-connCh:
			receiveMsg.Decode(receiveData)
			buf := fmt.Sprintf("%s: %s", receiveMsg.SourceName, string(receiveMsg.Body))
			fmt.Println(buf)
		}
	}
}
