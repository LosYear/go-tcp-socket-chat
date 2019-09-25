package server

import "net"

func (server *Server) hasUser(username string) bool {
	for _, user := range server.users {
		if user.username == username {
			return true
		}
	}

	return false
}

func (server Server) getUsername(connection net.Conn) string {
	value, found := server.users[connection.RemoteAddr().String()]

	if found {
		return value.username
	}

	return "Anonymous"
}

func (server *Server) performLogin(connection net.Conn, username string) bool {
	if server.hasUser(username) {
		return false
	}

	server.users[connection.RemoteAddr().String()] = User{username: username, connection: &connection}
	return true
}

func (server Server) usersList() []string {
	users := make([]string, 0, len(server.users))

	for _, user := range server.users {
		users = append(users, user.username)
	}

	return users
}
