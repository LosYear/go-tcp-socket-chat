package main

import (
	"./server"
	"fmt"
)

func main() {
	chatServer := server.InitServer(":1234")

	fmt.Println("Server listening on port", chatServer.Address)

	chatServer.Run()

}
