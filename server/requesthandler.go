package server

import (
	"net"
)

func (server Server) requestHandler(conn net.Conn, request Request) Response {
	if request.Name == LoginActionName {
		loggedIn := server.performLogin(conn, request.Payload)

		return Response{
			Name:  request.Name,
			Error: !loggedIn,
		}
	} else if request.Name == GetUsersActionName {
		return Response{
			Name:    request.Name,
			Payload: server.usersList(),
			Error:   false,
		}

	} else if request.Name == GetUsersCountActionName {
		return Response{
			Name:    request.Name,
			Payload: len(server.users),
			Error:   false,
		}

	} else if request.Name == SendMessageActionName {
		server.sendMessage(conn, request)

		return Response{
			Name:  "sent",
			Error: false,
		}

	} else if request.Name == LogoutActionName {

	}

	return Response{Name: "unknown", Message: "Unsupported command", Error: true}
}
