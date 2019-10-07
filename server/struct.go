package server

import "net"

type Server struct {
	Address string

	listener *net.TCPListener
	users    map[string]User
}

type User struct {
	username   string
	connection *net.Conn
}
