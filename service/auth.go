package service

import (
	"willow/response"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Login() response.Response {

	return response.Response{}
}

func (u *User) Create() response.Response {

	return response.Response{}
}
