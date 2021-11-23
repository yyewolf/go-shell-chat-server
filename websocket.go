package main

import (
	"encoding/json"
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
	go srv.ListenAndServe()
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
		_, message, err := c.ReadMessage()
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				currentUser.ConnectionClosed = true
				if currentUser.Username != "" {
					// Player has disconnected here

				}
			}
			break
		}
		currentUser.handleMessages(message)
	}
}

// handleMessages handles messages coming from the client
func (u *User) handleMessages(data []byte) {
	m := &Message{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return
	}
	d, err := json.Marshal(m.Data)
	if err != nil {
		return
	}
	switch m.Op {
	case identifyOp:
		packet := &CreateIdentify{}
		err = json.Unmarshal(d, &d)
		if err != nil {
			return
		}
		packet.Handle()
	case receiveMessageOp:
		packet := &CreateMessage{}
		err = json.Unmarshal(d, &d)
		if err != nil {
			return
		}
		packet.Handle()
	}
}
