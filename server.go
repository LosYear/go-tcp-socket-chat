package main

import (
	"./server"
	"fmt"
)

func main() {
	chatServer := server.InitServer(":64123")

	fmt.Println("Server listening on port", chatServer.Address)

	chatServer.Run()

}
