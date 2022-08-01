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
	fmt.Printf("Peer's name: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	peer := input.Text()
	fmt.Printf("Chat with peer: %s\n", peer)

	var connCh chan []byte = make(chan []byte)
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
		default:
			receiveMsg.Body = []byte{}
		}

		if len(receiveMsg.Body) > 0 {
			buf := fmt.Sprintf("%s: %s", receiveMsg.Name, string(receiveMsg.Body))
			fmt.Println(buf)
		}
	}
}
