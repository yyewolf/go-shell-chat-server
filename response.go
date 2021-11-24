package main

// Used to send a response to a user.
func (u *User) sendResponse(op int, v interface{}) {
	u.Connection.WriteJSON(&Message{
		Op:   op,
		Data: v,
	})
}

// Used to send a response to a user that sent a identify message.
func (u *User) identifyResponse(code int) {
	u.sendResponse(identifyOp, SendIdentify{
		Code: code,
	})
}

// Used to respond to a send message to a user.
func (u *User) receiveMessageResponse(code int) {
	u.sendResponse(receiveMessageOp, SendIdentify{
		Code: code,
	})
}

// Used to respond to a send file to a user.
func (u *User) receiveFileResponse(code int) {
	u.sendResponse(receiveFileOp, SendIdentify{
		Code: code,
	})
}

// Used to send a classic message to a user.
func (u *User) sendMessageClassic(message string) {
	u.sendResponse(sendMessageOp, SendMessage{
		Type:    messageClassic,
		User:    u.Username,
		Message: message,
	})
}

// Used to send a multiline message to a user.
func (u *User) sendMessageMultiline(messages []string) {
	u.sendResponse(sendMessageOp, SendMessage{
		Type:     messageClassic,
		User:     u.Username,
		Messages: messages,
	})
}

// Used to send a connect message to a user.
func (u *User) sendMessageConnect() {
	u.sendResponse(sendMessageOp, SendMessage{
		Type: messageConnection,
		User: u.Username,
	})
}

// Used to send a disconnect message to a user.
func (u *User) sendMessageDisconnect() {
	u.sendResponse(sendMessageOp, SendMessage{
		Type: messageDisconnection,
		User: u.Username,
	})
}
