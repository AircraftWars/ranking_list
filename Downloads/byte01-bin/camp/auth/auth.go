package auth

import (
	"bin/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var db = tools.GetDB()
var rdb = tools.GetRedisDB()

func LoginHandler(c *gin.Context) {
	request := LoginRequest{}
	if c.ShouldBind(&request) == nil {
		code := OK
		user := User{}
		// 根据Username查询相应user
		err := db.Where("user_name = ?", request.Username).Take(&user).Error
		if err != nil {
			fmt.Println(err)
			code = WrongPassword
			c.JSON(200, LoginResponse{
				Code: code,
				Data: struct{ UserID string }{UserID: user.UkUsername},
			})
			return
		} else {
			if user.Status == 0 {
				// 用户状态为已被删除
				code = UserHasDeleted
				c.JSON(200, LoginResponse{
					Code: code,
					Data: struct{ UserID string }{UserID: user.UkUsername},
				})
				return
			}
			// token格式为: username:userType
			token, err := tools.CreateToken(fmt.Sprintf("%s:%d", user.UkUsername, user.UserType))
			if err != nil {
				fmt.Println(err)
			} else {
				// 存入redis，没说存多久，就永久保存
				rdb.Set(user.UkUsername, token, 0)
				c.SetCookie("camp-session", token, 3600, "/", "127.0.0.1", false, true)
			}
		}
		c.JSON(200, LoginResponse{
			Code: code,
			Data: struct{ UserID string }{UserID: strconv.FormatInt(user.PkUserId, 10)},
		})
	} else {
		code := UnknownError
		c.JSON(200, LoginResponse{
			Code: code,
			Data: struct{ UserID string }{UserID: ""},
		})
	}

}

func LogoutHandler(c *gin.Context) {
	token, ckErr := c.Cookie("camp-session")
	code := LoginRequired
	if ckErr != nil {
		c.JSON(200, LogoutResponse{
			Code: code,
		})
		return
	}
	// 删除cookie
	c.SetCookie("camp-session", "", -1, "/", "127.0.0.1", false, true)
	userinfo, err := tools.ParseToken(token)
	if err != nil {
		code = UnknownError
		c.JSON(200, LogoutResponse{
			Code: code,
		})
		return
	}
	username := strings.Split(userinfo, ":")[0]
	// 删除redis中token
	rdb.Del(username)
	code = OK
	c.JSON(200, LogoutResponse{
		Code: code,
	})
}

func WhoAmIHandler(c *gin.Context) {
	code := OK
	// 检查cookie
	token, err := c.Cookie("camp-session")
	if err != nil {
		code = LoginRequired
		c.JSON(200, WhoAmIResponse{
			Code: code,
		})
		return
	}
	userinfo, err := tools.ParseToken(token)
	username := strings.Split(userinfo, ":")[0]
	user := User{}
	db.Where("user_name=?", username).Take(&user)
	c.JSON(200, WhoAmIResponse{
		Code: code,
		Data: struct {
			UserID   string
			Nickname string
			Username string
			UserType UserType
		}{UserID: strconv.FormatInt(user.PkUserId, 10), Nickname: user.Nickname, Username: username, UserType: UserType(user.UserType)},
	})

}
