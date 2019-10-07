package gui

import "github.com/losyear/go-tcp-socket-chat/client"

type State struct {
	IsLoggedIn bool
	Login      string
	Client     *client.Client
}

func NewState() *State {
	state := State{IsLoggedIn: false, Login: ""}

	return &state
}
