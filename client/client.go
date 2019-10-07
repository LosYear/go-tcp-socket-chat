package client

import (
	"encoding/json"
	"errors"
	"github.com/losyear/go-tcp-socket-chat/shared"
	"log"
	"net"
)

type Client struct {
	conn          net.Conn
	responseQueue []shared.Response
}

func InitClient(address string) *Client {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatal("Server is unreachable")
	}

	client := Client{conn: conn}

	return &client
}

func (c *Client) sendRequest(request shared.Request) error {
	jsonMessage, err := json.Marshal(request)

	if err != nil {
		return err
	}

	_, err = c.conn.Write([]byte(jsonMessage))

	return err
}

func (c *Client) awaitMessage() (shared.Response, error) {
	var buf [1024]byte
	bufSize, err := c.conn.Read(buf[0:])
	response := shared.Response{}

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(buf[:bufSize], &response)

	if err == nil {
		c.responseQueue = append(c.responseQueue, response)
	}

	return response, err

}

func (c *Client) Login(login string) error {
	loginRequest := shared.Request{Name: shared.LoginActionName, Payload: login}

	err := c.sendRequest(loginRequest)

	if err != nil {
		return err
	}

	response, err := c.awaitMessage()

	if err != nil {
		return err
	}

	if response.Error {
		return errors.New("user already exists")
	}

	return nil
}

func (c *Client) SendMessage(text string) error {
	sendRequest := shared.Request{Name: shared.SendMessageActionName, Payload: text}
	return c.sendRequest(sendRequest)
}

func (c *Client) SubscribeOnNotifications(handler func(shared.Response)) {
	go c.subscriptionHandler(handler)
}

func (c *Client) subscriptionHandler(handler func(shared.Response)) {
	for {
		response, err := c.awaitMessage()

		if err != nil {
			panic(err)
		}

		handler(response)
	}
}
