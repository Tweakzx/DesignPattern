package facade

import "fmt"

type User struct {
	Name string
	Msg  string
}

type IUser interface {
	Login(username, code string) (*User, error)
	Register(username, code string) (*User, error)
}

type IUserFacade interface {
	LoginOrRegister(username, code string) error
}

type UserService struct{}

func (us *UserService) Login(username, code string) (*User, error) {
	//假设系统里只有一个用户
	if username == "user0" && code == "123456" {
		return &User{Name: username, Msg: fmt.Sprintf("%s login success", username)}, nil
	}
	return nil, nil
}

func (us *UserService) Register(username, code string) (*User, error) {
	//注册账户
	return &User{Name: username, Msg: fmt.Sprintf("%s register success", username)}, nil
}

func (us UserService) LoginOrRegister(username, code string) (*User, error) {
	user, err := us.Login(username, code)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return us.Register(username, code)
}
