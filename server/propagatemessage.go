package server

import (
	"encoding/json"
	"github.com/losyear/go-tcp-socket-chat/shared"
)

func (server Server) propagateMessage(message shared.Response) {
	for _, user := range server.users {
		jsonMessage, _ := json.Marshal(message)

		_, err := (*user.connection).Write([]byte(jsonMessage))

		if err != nil {
			// todo: probably kick user?
			continue
		}
	}
}
