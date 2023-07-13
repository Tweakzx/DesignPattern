package proxy

import (
	"log"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

// @proxy IUser
type User struct {
}

func (u *User) Login(username, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (u *UserProxy) Login(username, password string) error {
	start := time.Now()

	if err := u.user.Login(username, password); err != nil {
		log.Printf("user login failed, err: %v\n", err)
		return err
	}

	log.Printf("user login success; cost time: %v", time.Since(start).Seconds())
	return nil
}
