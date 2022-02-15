package members

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//	/member/create
func Create(c *gin.Context) {
	member, code := &createMemberRequest{}, OK
	returnFunc := func(s string) {
		ret := createMemberResponse{
			Code: code,
			Data: struct{ UserID string }{UserID: s},
		}
		c.JSON(http.StatusOK, ret)
	}
	if err := c.Bind(member); err != nil {
		code = UnknownError
		returnFunc("")
		return
	}

	code = checkLogin(member.Username)
	if code != OK {
		returnFunc("")
		return
	}

	if !checkCreateParam(member) {
		code = ParamInvalid
	} else if checkUserHasExisted(member.Username) {
		code = UserHasExisted
	}
	if code != OK {
		returnFunc("")
	} else if createMember(member) == nil {
		userID := strconv.FormatInt(getCounts()-1, 10)
		returnFunc(userID)
	} else {
		code = UnknownError
		returnFunc("")
	}
}

//	/member
func GetOne(c *gin.Context) {
	userID, code := c.Query("UserID"), OK
	returnFunc := func(id int64, nickname, username string, userType UserType) {
		ret := getMemberResponse{
			Code: code,
			Data: TMember{
				UserID:   strconv.FormatInt(id, 10),
				Nickname: nickname,
				Username: username,
				UserType: userType,
			},
		}
		c.JSON(http.StatusOK, ret)
		return
	}
	returnNone := func() {
		ret := getMemberResponse{
			Code: code,
			Data: TMember{},
		}
		c.JSON(http.StatusOK, ret)
		return
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		code = ParamInvalid
		returnNone()
	}
	type User struct {
		UserID   int64
		Nickname string
		UserName string
		UserType UserType
		Status   int
	}
	user := &User{}
	result := db.Where("user_id = ?", id).First(&user)
	if err := result.Error; err != nil {
		code = UserNotExisted
	} else if user.Status == 0 {
		code = UserHasDeleted
	}
	if code != OK {
		returnNone()
	} else {
		returnFunc(user.UserID, user.Nickname, user.UserName, user.UserType)
	}
}

//	/member/list
func Gets(c *gin.Context) {
	Offset, Limit, code := c.Query("Offset"), c.Query("Limit"), OK
	returnNone := func() {
		ret := getMemberListResponse{
			Code: code,
			Data: struct {
				MemberList []TMember
			}{},
		}
		c.JSON(http.StatusOK, ret)
	}

	offset, err1 := strconv.Atoi(Offset)
	limit, err2 := strconv.Atoi(Limit)
	if err1 != nil || err2 != nil {
		code = ParamInvalid
		returnNone()
		return
	}

	var us []TMember
	result := db.Table("users").Where("status = ?", 1).Offset(offset).Limit(limit).Find(&us)
	if err := result.Error; err != nil {
		code = UnknownError
		returnNone()
		return
	}
	ret := getMemberListResponse{
		Code: code,
		Data: struct{ MemberList []TMember }{MemberList: us},
	}
	c.JSON(http.StatusOK, ret)
}

//	/member/update
func Update(c *gin.Context) {
	get, code := &updateMemberRequest{}, OK
	returnFunc := func() {
		ret := updateMemberResponse{
			code,
		}
		c.JSON(http.StatusOK, ret)
	}

	if err := c.Bind(get); err != nil {
		code = UnknownError
		returnFunc()
		return
	}

	userID, err := strconv.ParseInt(get.UserID, 10, 64)
	if err != nil || !checkNickname(get.Nickname) {
		code = ParamInvalid
		returnFunc()
		return
	}

	if !checkUserHasExistedById(userID) {
		code = UserNotExisted
		returnFunc()
		return
	}

	db.Model(&user{}).Where("user_id = ?", userID).Update("nickname", get.Nickname)
	returnFunc()
}

//	/member/delete
func Delete(c *gin.Context) {
	get, code := &deleteMemberRequest{}, OK
	returnFunc := func() {
		ret := deleteMemberResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, ret)
	}

	if err := c.Bind(get); err != nil {
		code = UnknownError
		returnFunc()
		return
	}

	id, err := strconv.ParseInt(get.UserID, 10, 64)
	if err != nil {
		code = ParamInvalid
		returnFunc()
		return
	}

	result := db.Model(&user{}).Where("user_id = ?", id).Update("status", 0)
	if err := result.Error; err != nil {
		code = UserNotExisted
	}
	returnFunc()
}
