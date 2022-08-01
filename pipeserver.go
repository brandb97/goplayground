package main

import (
	"fmt"
	"goplay/database"
	"goplay/message"
	"log"
	"net"
	"time"
)

var userDatabase database.DB

func server() {
	ln, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalln(err)
	}
	userDatabase = database.MakeDatabase()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	const MAXINPUT = 1024
	var buf []byte = make([]byte, MAXINPUT)
	var msg message.SendMsg

	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	msg.Decode(buf[:n])
	userDatabase.AddUser(msg.Name, &conn)
	fmt.Printf("%s log in\n", msg.Name)

	for err == nil {
		for !userDatabase.HasUser(msg.TargetName) {
			time.Sleep(50 * time.Millisecond)
		}
		targetConn := *userDatabase.UsertoConn(msg.TargetName)
		rmsg := message.ReceiveMsg{
			Name:       msg.TargetName,
			SourceName: msg.Name,
			Body:       msg.Body,
		}
		rbuf := rmsg.Encode()
		_, err = targetConn.Write(rbuf)

		n, err = conn.Read(buf)
		msg.Decode(buf[:n])
	}
	userDatabase.DeleteUser(msg.Name)
}
