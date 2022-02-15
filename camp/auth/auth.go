package auth

import (
	"bin/tools"
	"bin/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var db = tools.GetDB()
var rdb = tools.GetRedisDB()

func LoginHandler(c *gin.Context) {
	request := types.LoginRequest{}
	if c.ShouldBind(&request) == nil {
		code := types.OK
		user := User{}
		// 根据Username查询相应user
		err := db.Where("uk_username = ?", request.Username).Take(&user).Error
		if err != nil {
			fmt.Println(err)
			code = types.WrongPassword
			c.JSON(200, types.LoginResponse{
				Code: code,
				Data: struct{ UserID string }{UserID: user.UkUsername},
			})
			return
		} else {
			// token格式为: username:userType
			token, err := tools.CreateToken(fmt.Sprintf("%s:%d", user.UkUsername, user.UserType))
			if err != nil {
				fmt.Println(err)
			} else {
				// 存入redis，没说存多久，就永久保存
				rdb.Set(user.UkUsername, token, 0)
				c.SetCookie("camp-session", token, 3600, "/", "localhost", false, true)
			}
		}
		c.JSON(200, types.LoginResponse{
			Code: code,
			Data: struct{ UserID string }{UserID: user.UkUsername},
		})
	} else {
		//fmt.Println("nil!")
	}

}

func LogoutHandler(c *gin.Context) {
	token, _ := c.Cookie("camp-session")
	// 删除cookie
	c.SetCookie("camp-session", "", -1, "/", "localhost", false, true)
	userinfo, err := tools.ParseToken(token)
	if err != nil {
		return
	}
	username := strings.Split(userinfo, ":")[0]
	// 删除redis中token
	rdb.Del(username)
	code := types.OK
	c.JSON(200, types.LogoutResponse{
		Code: code,
	})
}

func WhoAmIHandler(c *gin.Context) {
	code := types.OK
	// 检查cookie
	token, err := c.Cookie("camp-session")
	if err != nil {
		code = types.LoginRequired
		c.JSON(200, types.WhoAmIResponse{
			Code: code,
		})
		return
	}
	userinfo, err := tools.ParseToken(token)
	username := strings.Split(userinfo, ":")[0]
	user := User{}
	db.Where("uk_username=?", username).Take(&user)
	c.JSON(200, types.WhoAmIResponse{
		Code: code,
		Data: struct {
			UserID   string
			Nickname string
			Username string
			UserType types.UserType
		}{UserID: strconv.FormatInt(user.PkUserId, 10), Nickname: user.Nickname, Username: username, UserType: types.UserType(user.UserType)},
	})

}
