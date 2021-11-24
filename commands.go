package main

import "fmt"

func commandLoader() {
	commands = make(map[string]func(*commandCtx))
	commands["ls"] = sendList
	commands["@"] = dm
}

func sendList(c *commandCtx) {
	var str []string
	str = append(str, "List of users connected :")
	for _, u := range connections {
		str = append(str, "   - "+u.Username)
	}
	c.User.sendResponse(sendMessageOp, SendMessage{
		Type:     messageMultiline,
		User:     "Serveur",
		Messages: str,
	})
}

func dm(c *commandCtx) {
	if len(c.Args) < 2 {
		c.User.sendResponse(sendMessageOp, SendMessage{
			Type:    messageClassic,
			User:    "Serveur",
			Message: fmt.Sprintf("Pas de message."),
		})
		return
	}
	user := c.Args[0]
	msgS := c.Args[1:]
	var msg string
	for _, m := range msgS {
		msg += m + " "
	}
	target, found := connections.Get(user)
	if !found {
		c.User.sendResponse(sendMessageOp, SendMessage{
			Type:    messageClassic,
			User:    "Serveur",
			Message: fmt.Sprintf("Utilisateur %s non trouvé.", user),
		})
		return
	}
	target.sendResponse(sendMessageOp, SendMessage{
		Type:    messageDM,
		User:    c.User.Username,
		Message: msg,
	})
}
