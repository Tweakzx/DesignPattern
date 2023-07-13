package mediator

import "fmt"

type Mediator interface {
	Send(message string, sender User)
	Register(user User)
}

type ChatRoom struct {
	roomName string
	users    map[string]User
}

func CreateChatRoom(name string) *ChatRoom {
	return &ChatRoom{name, make(map[string]User)}
}

func (c *ChatRoom) Send(message string, sender User) {
	for _, user := range c.users {
		if user.name != sender.name {
			user.Receive(message, sender.name, c.roomName)
		}
	}
}

func (c *ChatRoom) Register(user User) {
	c.users[user.name] = user
}

type User struct {
	name     string
	age      int
	sex      string
	mediator Mediator
}

func (u *User) EnterRoom(mediator Mediator) {
	u.mediator = mediator
	u.mediator.Register(*u)
	u.mediator.Send("User "+u.name+" has entered the room", *u)
}

func (u *User) Send(message string) {
	u.mediator.Send(message, *u)
}

func (u *User) Receive(message string, from string, at string) {
	fmt.Println(from + ": " + message)
}
