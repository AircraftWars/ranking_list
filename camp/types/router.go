package types

import (
	"bin/auth"
	"bin/members"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create", members.Create)
	g.GET("/member", members.GetOne)
	g.GET("/member/list", members.Gets)
	g.POST("/member/update", members.Update)
	g.POST("/member/delete", members.Delete)

	// 登录
	g.POST("/auth/login", auth.LoginHandler)
	g.POST("/auth/logout", auth.LogoutHandler)
	g.GET("/auth/whoami", auth.WhoAmIHandler)
	/*
		// 排课
		g.POST("/course/create")
		g.GET("/course/get")

		g.POST("/teacher/bind_course")
		g.POST("/teacher/unbind_course")
		g.GET("/teacher/get_course")
		g.POST("/course/schedule")

		// 抢课
		g.POST("/student/book_course")
		g.GET("/student/course")
	*/
}