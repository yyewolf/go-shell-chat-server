package main

// Gets a user from a pool
func (p *Pool) Get(username string) (user *User, found bool) {
	s := *p
	for _, u := range s {
		if u.Username == username {
			return u, true
		}
	}
	return nil, false
}

// Check if a users is in a pool
func (p *Pool) Exists(username string) bool {
	s := *p
	for _, u := range s {
		if u.Username == username {
			return true
		}
	}
	return false
}

// Removes a user from a pool
func (p *Pool) Remove(username string) bool {
	s := *p
	for i, u := range s {
		if u.Username == username {
			s = append(s[:i], s[i+1:]...)
			*p = s
			return true
		}
	}
	return false
}

// Broadcast a message to the pool
func (p *Pool) Broadcast(op int, v interface{}) {
	s := *p
	for _, u := range s {
		if u.ConnectionClosed {
			continue
		}
		u.Connection.WriteJSON(Message{
			Op:   op,
			Data: v,
		})
	}
}
