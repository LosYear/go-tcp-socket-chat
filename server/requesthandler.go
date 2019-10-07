package server

import (
	"github.com/losyear/go-tcp-socket-chat/shared"
	"net"
)

func (server Server) requestHandler(conn net.Conn, request shared.Request) *shared.Response {
	if request.Name == shared.LoginActionName {
		loggedIn := server.performLogin(conn, request.Payload)

		return &shared.Response{
			Name:  request.Name,
			Error: !loggedIn,
		}
	} else if request.Name == shared.GetUsersActionName {
		return &shared.Response{
			Name:    request.Name,
			Payload: server.usersList(),
			Error:   false,
		}

	} else if request.Name == shared.GetUsersCountActionName {
		return &shared.Response{
			Name:    request.Name,
			Payload: len(server.users),
			Error:   false,
		}

	} else if request.Name == shared.SendMessageActionName {
		server.sendMessage(conn, request)

		return nil

	} else if request.Name == shared.LogoutActionName {

	}

	return &shared.Response{Name: "unknown", Message: "Unsupported command", Error: true}
}
