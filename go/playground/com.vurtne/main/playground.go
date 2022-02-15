package main

import (
	_ "crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"playground/com.vurtne/auth"
	"playground/com.vurtne/tools"
)

type User struct {
	//gorm.Model
	//id int64 `gorm:column:id; PRIMARY_KEY`
	id int64
	//Name     string `gorm:"column:name"`
	Name     string
	Password string `gorm:"column:password"`
	//Updated   int64 `gorm:"autoUpdateTime:nano"` // 自定义字段， 使用时间戳填纳秒数充更新时间
	//Updated   int64 `gorm:"autoUpdateTime:milli"` //自定义字段， 使用时间戳毫秒数填充更新时间
	//Created   int64 `gorm:"autoCreateTime"`      //自定义字段， 使用时间戳秒数填充创建时间
}

func (u User) TableName() string {
	// 表名
	return "test01"
}

var db = tools.GetDB()
var rdb = tools.GetRedisDB()

func main() {
	//获取db
	//db := tools.GetDB()

	//u := User{
	//	Name:     "kevin",
	//	Password: "123321",
	//}
	//if err := db.Create(&u).Error; err != nil {
	//	fmt.Println("插入失败", err)
	//	return
	//}

	//u := User{}
	////result := db.Where("name = ?", "kevin").First(&u)
	//// 使用Debug()打印日志
	//result := db.Debug().Where("name = ?", "kevin").First(&u)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	fmt.Println("找不到记录")
	//	return
	//}
	//
	//var users []User
	//db.Debug().Find(&users)
	//fmt.Println(users)

	router := gin.Default()
	router.GET("/cookie", func(context *gin.Context) {
		context.SetCookie("abc", "value", 3600, "/", "localhost", false, true)
	})
	router.GET("/getCookie", func(context *gin.Context) {
		Handler(context)
	})
	router.GET("/redisTest", func(context *gin.Context) {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_, err := rdb.Ping().Result()
		if err != nil {
			fmt.Println("error to connect redis!")
		}
		fmt.Println("success")
		rdb.Set("goredis", 100, 0)
	})

	router.POST("/loginTest", func(context *gin.Context) {
		auth.LoginHandler(context)
	})

	router.POST("/logoutTest", func(context *gin.Context) {
		auth.LogoutHandler(context)
	})

	router.POST("/whoamiTest", func(context *gin.Context) {
		auth.WhoAmIHandler(context)
	})

	router.Run()

}
func Handler(c *gin.Context) {
	data, err := c.Cookie("abc")
	if err != nil {
		c.String(404, "not found!")
		return
	}
	c.String(200, data)

}
