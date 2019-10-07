package server

import (
	"encoding/json"
	"github.com/losyear/go-tcp-socket-chat/shared"
	"log"
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

		log.Println("Request:", string(buf[0:]))

		request := shared.Request{}

		err = json.Unmarshal(buf[:bufSize], &request)

		if err != nil {
			return
		}

		result := server.requestHandler(conn, request)

		if result == nil {
			continue
		}

		response, _ := json.Marshal(result)

		_, err = conn.Write([]byte(response))

		if err != nil {
			return
		}

	}
}
