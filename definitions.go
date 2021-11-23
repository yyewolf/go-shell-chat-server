package main

import "github.com/gorilla/websocket"

// Defines OP names
const (
	identifyOp = iota
	receiveMessageOp
	sendMessageOp
)

// Defines messages type
const (
	messageClassic = iota
	messageConnection
	messageDisconnection
	messageMultiline
)

// Defines codes
const (
	codeSuccess = 200
	codeError   = 400
)

// Defines users
type User struct {
	Connection       *websocket.Conn `json:"-"`
	ConnectionClosed bool            `json:"-"`
	Username         string          `json:"username"`
}

// Defines special type for pool
type Pool []*User

// Defines global pool
var connections Pool

// Defines the messages standard
type Message struct {
	Op   int         `json:"op"`
	Data interface{} `json:"data"`
}

type CreateIdentify struct {
	Username string `json:"username"`
}

type CreateMessage struct {
	Type     string   `json:"type"`
	Message  string   `json:"message"`
	Messages []string `json:"messages"`
}

type Identify struct {
	Code int `json:"code"`
}

type SendMessage struct {
	Type     string   `json:"type"`
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Messages []string `json:"messages"`
}
