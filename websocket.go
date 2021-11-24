package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/gorilla/websocket"
)

// Upgrader for websocket (low buffer size for stability)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  64,
	WriteBufferSize: 64,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func startListening() {
	mux := pat.New()
	port := ":30"
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}
	mux.Get("/", http.HandlerFunc(websocketLoop))
	go func() {
		err := srv.ListenAndServe()
		fmt.Println(err)
	}()
}

func websocketLoop(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("err while upgrading:", err)
		return
	}
	defer c.Close()
	currentUser := &User{
		Connection: c,
	}
	for {
		if currentUser.ConnectionClosed {
			return
		}
		msg := &Message{}
		err := c.ReadJSON(msg)
		if err != nil {
			currentUser.ConnectionClosed = true
			if currentUser.Username != "" {
				connections.Remove(currentUser.Username)
				connections.Broadcast(sendMessageOp, SendMessage{
					Type: messageDisconnection,
					User: currentUser.Username,
				})
			}
			connections.Remove(currentUser.Username)
			return
		}
		currentUser.handleMessages(msg)
	}
}

// handleMessages handles messages coming from the client
func (u *User) handleMessages(m *Message) {
	d, err := json.Marshal(m.Data)
	if err != nil {
		return
	}
	switch m.Op {
	case identifyOp:
		packet := &CreateIdentify{}
		err = json.Unmarshal(d, &packet)
		if err != nil {
			return
		}
		packet.Handle(u)
	case receiveMessageOp:
		packet := &CreateMessage{}
		err = json.Unmarshal(d, &packet)
		if err != nil {
			return
		}
		packet.Handle(u)
	case receiveFileOp:
		packet := &File{}
		err = json.Unmarshal(d, &packet)
		if err != nil {
			return
		}
		packet.Handle(u)
	}
}
