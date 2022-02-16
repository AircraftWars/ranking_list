package members

import (
	"bin/tools"
	"github.com/go-redis/redis"
)

//-------------
//数据库操纵函数
//-------------

var db = tools.GetDB()
var rdb = tools.GetRedisDB()

//创建用户
func createMember(member *createMemberRequest) error {
	instance := user{
		UserId:   getCounts(),
		UserName: member.Username,
		Nickname: member.Nickname,
		Password: member.Password,
		UserType: int(member.UserType),
		Status:   1,
	}
	result := db.Model(&user{}).Create(&instance)
	return result.Error
}

//判断用户是否存在by user_name
func checkUserHasExisted(s string) bool {
	var count int64
	db.Table("users").Where("user_name = ?", s).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

//判断用户是否存在by user_id
func checkUserHasExistedById(id int64) bool {
	var count int64
	db.Table("users").Where("user_id = ?", id).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

//获取users表总行数，作为UserID
func getCounts() int64 {
	var count int64
	db.Table("users").Count(&count)
	return count
}

//检验是否登录 及 是否具有管理员权限
func checkLogin(s string) ErrNo {
	token := rdb.Get(s)
	if token.Err() == redis.Nil {
		return LoginRequired
	}
	t := token.Val()
	if t[len(t)-1] == '1' {
		return PermDenied
	}
	return OK
}
