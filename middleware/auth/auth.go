package auth

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"willow/config"
	"willow/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, response.Error(response.ErrAuth))
			c.Abort()
			return
		}
		// if service.IsBlacklist(token) {
		// 	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		// 	c.Abort()
		// 	return
		// }
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, response.Error(response.ErrAuthExpired))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, response.ErrorUnknown(response.ErrAuthUnknown, err.Error()))
			c.Abort()
			return
		}
		// if err, _ = service.FindUserByUuid(claims.UUID.String()); err != nil {
		// 	_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
		// 	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		// 	c.Abort()
		// }

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + int64(config.Conf.Jwt.Expired)
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			// if global.GVA_CONFIG.System.UseMultipoint {
			// 	err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
			// 	if err != nil {
			// 		global.GVA_LOG.Error("get redis jwt failed", zap.Any("err", err))
			// 	} else { // 当之前的取成功时才进行拉黑操作
			// 		_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
			// 	}
			// 	// 无论如何都要记录当前的活跃状态
			// 	_ = service.SetRedisJWT(newToken, newClaims.Username)
			// })
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Conf.Jwt.SigningKey),
	}
}

type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {

	prefix := "Bearer "
	if tokenString != "" && strings.HasPrefix(tokenString, prefix) {
		tokenString = tokenString[len(prefix):]
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
