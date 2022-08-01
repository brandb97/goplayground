package main

import (
	"bufio"
	"fmt"
	"goplay/message"
	"log"
	"net"
	"os"
	"time"
)

const serverIp = "127.0.0.1"
const serverPort = ":1234"

func client() {
	conn, err := net.Dial("tcp", serverIp+serverPort)
	if err != nil {
		log.Fatalln(err)
	}
	name := MakeName()
	fmt.Printf("login %s%s using name: %s\n", serverIp, serverPort, name)
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
	input.Scan()
	peer := input.Text()
	fmt.Printf("Chat with peer: %s\n", peer)

	var connCh chan []byte
	go func() {
		const MAXINPUT = 1024
		var data []byte = make([]byte, MAXINPUT)
		_, err := conn.Read(data)
		if err != nil {
			connCh <- data
		} else {
			close(connCh)
			log.Fatal(err)
		}
	}()

	sendMsg := message.SendMsg{
		Name:       name,
		TargetName: peer,
	}
	for input.Scan() {
		sendMsg.Body = []byte(input.Text())
		data := sendMsg.Encode()
		_, err := conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}

		var receiveMsg message.ReceiveMsg
		select {
		case data = <-connCh:
			receiveMsg.Decode(data)
		case <-time.After(50 * time.Millisecond):
			receiveMsg.Body = []byte{}
		}

		if len(receiveMsg.Body) > 0 {
			buf := fmt.Sprintf("%s: %s", receiveMsg.Name, string(receiveMsg.Body))
			fmt.Println(buf)
		}
	}
}
