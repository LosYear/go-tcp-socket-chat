package serverapp

import (
	"fmt"
	"github.com/losyear/go-tcp-socket-chat/server"
	"os"
)

func main() {
	address := ":64123"

	if len(os.Args) > 1 && os.Args[1] != "" {
		address = os.Args[1]
	}

	chatServer := server.InitServer(address)

	fmt.Println("Server listening on port", chatServer.Address)

	chatServer.Run()

}
