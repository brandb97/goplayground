package main

import (
	"goplay/database"
	"goplay/message"
	"log"
	"net"
	"time"
)

var userDatabase database.DB

func server() {
	ln, err := net.Listen("tcp", ":8000")
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
	const MAXBUF = 1024
	var buf []byte = make([]byte, MAXBUF)
	var msg message.SendMsg

	_, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	msg.Decode(buf)
	userDatabase.AddUser(msg.Name, &conn)
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

		_, err = conn.Read(buf)
		msg.Decode(buf)
	}
	userDatabase.DeleteUser(msg.Name)
}
