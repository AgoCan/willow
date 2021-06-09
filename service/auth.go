package service

import (
	"database/sql"
	"errors"
	"fmt"
	"willow/global"
	"willow/model"
	"willow/response"
	"willow/utils/hash"

	"gorm.io/gorm"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	Nickname string `json:"nickName"`
}

func (u *UserLogin) Login() response.Response {

	return response.Response{}
}

func (r *Register) Create() response.Response {
	var u model.User
	u.Username = r.Username
	u.Password = r.Password
	u.Nickname = sql.NullString{Valid: true, String: r.Nickname}

	if !errors.Is(global.GDB.Where("username = ?", u.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.UserExist)
	}

	u.Password = hash.MD5V([]byte(r.Password))
	fmt.Println(u.Password)
	err := global.GDB.Create(&u).Error
	if err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Response{}
}
