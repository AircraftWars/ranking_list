package members

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1  // 参数不合法
	UserHasExisted ErrNo = 2  // 该 Username 已存在
	UserHasDeleted ErrNo = 3  // 用户已删除
	UserNotExisted ErrNo = 4  // 用户不存在
	LoginRequired  ErrNo = 6  // 用户未登录
	PermDenied     ErrNo = 10 // 没有操作权限

	UnknownError ErrNo = 255 // 未知错误
)

type TMember struct {
	UserID   string   `gorm:"column:user_id"`
	Nickname string   `gorm:"column:nickname"`
	Username string   `gorm:"column:user_name"`
	UserType UserType `gorm:"column:user_type"`
}

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

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
