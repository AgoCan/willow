package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"willow/config"
	"willow/global"
	"willow/middleware/auth"
	"willow/model"
	"willow/response"
	"willow/utils/hash"

	"github.com/dgrijalva/jwt-go"
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
	u.Password = hash.MD5V([]byte(u.Password))
	var modelU model.User
	modelU.Username = u.Username
	modelU.Password = u.Password
	err := global.GDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&modelU).Error
	if err != nil {
		return response.Error(response.ErrUsernameOrPassword)
	}
	j := &auth.JWT{SigningKey: []byte(config.Conf.Jwt.SigningKey)} // 唯一签名
	claims := auth.CustomClaims{
		ID:       modelU.ID,
		NickName: modelU.Nickname.String,
		Username: modelU.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                           // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(config.Conf.Jwt.Expired), // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                           // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return response.Error(response.ErrJwtToken)
	}
	modelU.Token = token
	return response.Success(modelU)
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
	return response.Success("成功创建用户")
}
