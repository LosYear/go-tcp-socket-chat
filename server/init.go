package server

import "net"

func InitServer(address string) *Server {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	server := Server{Address: address, listener: listener, users: make(map[string]User)}

	return &server
}
