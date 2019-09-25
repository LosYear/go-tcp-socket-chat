package server

import "log"

func (server Server) Run() {
	for {
		conn, err := server.listener.Accept()

		if err != nil {
			return
		}

		log.Println("Accepted new client", conn.RemoteAddr())

		go server.clientHandler(conn)
	}
}
