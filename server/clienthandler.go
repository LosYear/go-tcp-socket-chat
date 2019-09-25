package server

import (
	"encoding/json"
	"net"
)

func (server Server) clientHandler(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [1024]byte

		bufSize, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		request := Request{}

		err = json.Unmarshal(buf[:bufSize], &request)

		if err != nil {
			return
		}

		response, _ := json.Marshal(server.requestHandler(conn, request))

		_, err = conn.Write([]byte(response))

		if err != nil {
			return
		}

	}
}
