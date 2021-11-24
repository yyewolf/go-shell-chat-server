package main

import "github.com/gorilla/websocket"

// Defines OP names
const (
	identifyOp = iota
	receiveMessageOp
	sendMessageOp
	receiveFileOp
	sendFileOp
)

// Defines messages type
const (
	messageClassic = iota
	messageConnection
	messageDisconnection
	messageMultiline
	messageDM
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
	Data interface{} `json:"data,omitempty"`
}

type CreateIdentify struct {
	Username string `json:"username,omitempty"`
}

type CreateMessage struct {
	Type     int      `json:"type"`
	Message  string   `json:"message,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type SendIdentify struct {
	Code int `json:"code,omitempty"`
}

type SendMessage struct {
	Type     int      `json:"type"`
	User     string   `json:"user,omitempty"`
	Message  string   `json:"message,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type File struct {
	Name string `json:"name"`
	User string `json:"user,omitempty"`
	Data []byte `json:"data,omitempty"`
}

// Server side commands
var commands map[string]func(*commandCtx)

type commandCtx struct {
	Args []string
	User *User
}
