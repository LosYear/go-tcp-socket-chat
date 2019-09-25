package server

import "encoding/json"

func (server Server) propagateMessage(message Response) {
	for _, user := range server.users {
		jsonMessage, _ := json.Marshal(message)

		_, err := (*user.connection).Write([]byte(jsonMessage))

		if err != nil {
			// todo: probably kick user?
			continue
		}
	}
}
