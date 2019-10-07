package main

import (
	"fmt"
	"github.com/losyear/go-tcp-socket-chat/server"
)

func main() {
	chatServer := server.InitServer(":64123")

	fmt.Println("Server listening on port", chatServer.Address)

	chatServer.Run()

}
