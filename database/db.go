package database

import (
	"net"
)

type DB map[string]*net.Conn

func MakeDatabase() DB {
	return DB{}
}

func (db DB) AddUser(name string, conn *net.Conn) {
	db[name] = conn
}

func (db DB) DeleteUser(name string) {
	db[name] = nil
}

func (db DB) HasUser(name string) bool {
	return db[name] != nil
}

func (db DB) UsertoConn(name string) *net.Conn {
	return db[name]
}
