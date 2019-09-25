package server

import "net"

func (server Server) sendMessage(conn net.Conn, request Request) {
	server.propagateMessage(Response{
		Name:    NewMessageNotificationName,
		Error:   false,
		Payload: TextMessage{Username: server.getUsername(conn), Text: request.Payload},
	})
}
