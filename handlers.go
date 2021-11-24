package main

import (
	"strings"
)

func (i *CreateIdentify) Handle(u *User) {
	if u.Username != "" {
		u.identifyResponse(codeError)
		return
	}
	if connections.Exists(i.Username) {
		u.identifyResponse(codeError)
		return
	}
	u.Username = i.Username
	connections.Broadcast(sendMessageOp, SendMessage{
		Type: messageConnection,
		User: u.Username,
	})
	connections = append(connections, u)
	u.identifyResponse(codeSuccess)
}

func (m *CreateMessage) Handle(u *User) {
	if u.Username == "" {
		u.receiveMessageResponse(codeError)
		return
	}

	if m.Type == messageClassic {
		if strings.HasPrefix(m.Message, commandPrefix) || strings.HasPrefix(m.Message, "@") {
			input := strings.Replace(m.Message, "/", "", 1)
			splt := strings.Split(input, " ")
			if len(splt) == 0 {
				return
			}
			command := splt[0]
			if strings.HasPrefix(command, "@") {
				splt = append([]string{splt[0], strings.Replace(command, "@", "", 1)}, splt[1:]...)
				command = "@"
			}
			command = strings.ToLower(command)
			call, found := commands[command]
			if !found {
				return
			}
			args := []string{}
			if len(splt) > 1 {
				args = splt[1:]
			}
			ctx := &commandCtx{
				Args: args,
				User: u,
			}
			call(ctx)
			return
		}
	}

	connections.Broadcast(sendMessageOp, SendMessage{
		Type:     m.Type,
		User:     u.Username,
		Message:  m.Message,
		Messages: m.Messages,
	})
	u.receiveMessageResponse(codeSuccess)
}

func (m *File) Handle(u *User) {
	if u.Username == "" {
		u.receiveFileResponse(codeError)
		return
	}

	target, found := connections.Get(m.User)
	if !found {
		u.receiveFileResponse(codeError)
		return
	}

	target.sendResponse(sendFileOp, File{
		Name: m.Name,
		Data: m.Data,
		User: u.Username,
	})
	u.receiveFileResponse(codeSuccess)
}
