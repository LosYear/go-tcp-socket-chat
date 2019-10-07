package server

import (
	"github.com/losyear/go-tcp-socket-chat/shared"
	"net"
)

func (server Server) sendMessage(conn net.Conn, request shared.Request) {
	server.propagateMessage(shared.Response{
		Name:    shared.NewMessageNotificationName,
		Error:   false,
		Payload: shared.TextMessage{Username: server.getUsername(conn), Text: request.Payload},
	})
}
