package main

func (i *CreateIdentify) Handle(u *User) {
	if u.Username != "" {
		u.identifyResponse(codeError)
		return
	}
	if connections.Exists(i.Username) {
		u.identifyResponse(codeError)
		return
	}
	connections.Broadcast(sendMessageOp, SendMessage{
		Type: messageConnection,
		User: u.Username,
	})
	connections = append(connections, u)
	u.Username = i.Username
	u.identifyResponse(codeSuccess)
}

func (m *CreateMessage) Handle(u *User) {
	if u.Username == "" {
		u.receiveMessageResponse(codeError)
		return
	}
	connections.Broadcast(sendMessageOp, SendMessage{
		Type:    messageClassic,
		User:    u.Username,
		Message: m.Message,
	})
	u.receiveMessageResponse(codeSuccess)
}
