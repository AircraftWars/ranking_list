package members

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

type ErrNo int

const (
	OK                 ErrNo = 0
	ParamInvalid       ErrNo = 1  // 参数不合法
	UserHasExisted     ErrNo = 2  // 该 Username 已存在
	UserHasDeleted     ErrNo = 3  // 用户已删除
	UserNotExisted     ErrNo = 4  // 用户不存在
	WrongPassword      ErrNo = 5  // 密码错误
	LoginRequired      ErrNo = 6  // 用户未登录
	CourseNotAvailable ErrNo = 7  // 课程已满
	CourseHasBound     ErrNo = 8  // 课程已绑定过
	CourseNotBind      ErrNo = 9  // 课程未绑定过
	PermDenied         ErrNo = 10 // 没有操作权限
	StudentNotExisted  ErrNo = 11 // 学生不存在
	CourseNotExisted   ErrNo = 12 // 课程不存在
	StudentHasNoCourse ErrNo = 13 // 学生没有课程
	StudentHasCourse   ErrNo = 14 // 学生有课程

	UnknownError ErrNo = 255 // 未知错误
)

type TMember struct {
	UserID   string   `gorm:"column:user_id"`
	Nickname string   `gorm:"column:nickname"`
	Username string   `gorm:"column:user_name"`
	UserType UserType `gorm:"column:user_type"`
}

// 成员管理

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加

type createMemberRequest struct {
	Nickname string   `form:"Nickname"` // required，不小于 4 位 不超过 20 位
	Username string   `form:"Username"` // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   `form:"Password"` // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType `form:"UserType"` // required, 枚举值
}

type createMemberResponse struct {
	Code ErrNo
	Data struct {
		UserID string // int64 范围
	}
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type getMemberResponse struct {
	Code ErrNo
	Data TMember
}

// 批量获取成员信息
type getMemberListResponse struct {
	Code ErrNo
	Data struct {
		MemberList []TMember
	}
}

// 更新成员信息

type updateMemberRequest struct {
	UserID   string
	Nickname string
}

type updateMemberResponse struct {
	Code ErrNo
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type deleteMemberRequest struct {
	UserID string
}

type deleteMemberResponse struct {
	Code ErrNo
}

type user struct {
	UserId   int64
	UserName string
	Nickname string
	Password string
	UserType int
	Status   int
}
