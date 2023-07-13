package mediator

import "testing"

func TestMediator(t *testing.T) {
	ChatRoom := CreateChatRoom("Golang Chat Room")

	John := &User{name: "John", sex: "male", age: 22}
	Jane := &User{name: "Jane", sex: "female", age: 20}
	Jack := &User{name: "Jack", sex: "male", age: 25}

	John.EnterRoom(ChatRoom)
	Jane.EnterRoom(ChatRoom)
	Jack.EnterRoom(ChatRoom)

	John.Send("Hi, I'm John.")
	Jane.Send("Hello, I'm Jane.")
	Jack.Send("Nice to meet you, I'm Jack.")
}
