package main

import (
	"github.com/losyear/go-tcp-socket-chat/client/gui"
	"os"
)

func main() {
	app := gui.NewApplication()

	address := "127.0.0.1:64123"

	if len(os.Args) > 1 && os.Args[1] != "" {
		address = os.Args[1]
	}

	app.Run(address)
}
