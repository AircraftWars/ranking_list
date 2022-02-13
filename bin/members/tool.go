package members

//----------
//工具函数
//----------

import (
	"github.com/gin-gonic/gin"
	"unicode"
)

//检查昵称是否合法
func checkNickname(s string) bool {
	n := len(s) //len()返回字节长度
	if n < 4 || n > 20 {
		return false
	}
	return true
}

//检查用户名是否合法
func checkUsername(s string) bool {
	n := len(s)
	if n < 8 || n > 20 {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

//检查密码是否合法
func checkPassword(s string) bool {
	n := len(s)
	if n < 8 || n > 20 {
		return false
	}
	var upper, lower, digit bool
	for _, r := range s {
		if unicode.IsUpper(r) {
			upper = true
		} else if unicode.IsLower(r) {
			lower = true
		} else if unicode.IsDigit(r) {
			digit = true
		} else {
			return false
		}
	}
	if upper && lower && digit {
		return true
	}
	return false
}

//检查用户类型是否合法
func checkUserType(tp UserType) bool {
	if tp == Admin || tp == Student || tp == Teacher {
		return true
	}
	return false
}

//	检查创建成员时的参数合法性
func checkCreateParam(m *createMemberRequest) bool {
	f := true
	f = f && checkNickname(m.Nickname)
	f = f && checkUsername(m.Username)
	f = f && checkPassword(m.Password)
	f = f && checkUserType(m.UserType)
	return f
}

//检验是否登录
func checkLogin(c *gin.Context) bool {
	/*
		-----------
	*/
	return true
}

//检验是否有管理员权限
func checkIsAdmin(c *gin.Context) bool {
	/*
		-----------
	*/
	return true
}
