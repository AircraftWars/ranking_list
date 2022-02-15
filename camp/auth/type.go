package auth

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1 // 参数不合法
	UserHasExisted ErrNo = 2 // 该 Username 已存在
	UserHasDeleted ErrNo = 3 // 用户已删除
	UserNotExisted ErrNo = 4 // 用户不存在
	WrongPassword  ErrNo = 5 // 密码错误
	LoginRequired  ErrNo = 6 // 用户未登录

	UnknownError ErrNo = 255 // 未知错误
)

type ResponseMeta struct {
	Code ErrNo
}

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

type TMember struct {
	UserID   string
	Nickname string
	Username string
	UserType UserType
}

type TCourse struct {
	CourseID  string
	Name      string
	TeacherID string
}

// ----------------------------------------
// 登录

type LoginRequest struct {
	Username string
	Password string
}

// 登录成功后需要 Set-Cookie("camp-session", ${value})
// 密码错误范围密码错误状态码

type LoginResponse struct {
	Code ErrNo
	Data struct {
		UserID string
	}
}

// 登出

type LogoutRequest struct{}

// 登出成功需要删除 Cookie

type LogoutResponse struct {
	Code ErrNo
}

// WhoAmI 接口，用来测试是否登录成功，只有此接口需要带上 Cookie

type WhoAmIRequest struct {
}

// 用户未登录请返回用户未登录状态码

type WhoAmIResponse struct {
	Code ErrNo
	Data TMember
}
