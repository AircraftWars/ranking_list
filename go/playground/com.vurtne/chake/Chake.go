package chake

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"playground/com.vurtne/entity"
	"playground/com.vurtne/types"
	"strconv"
)

var DB *gorm.DB

func Chake(c *gin.Context) {
	ID := c.Query("StudentID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	code := types.OK
	tcourse := make([]types.TCourse, 0, 8)
	course := make([]entity.Course, 0, 8)
	defer func() {
		c.JSON(200, types.GetStudentCourseResponse{
			Code: code,
			Data: struct {
				CourseList []types.TCourse
			}{tcourse},
		})
	}()

	var user entity.User
	if err := DB.Model(&entity.User{}).Where("pk_user_id = ? and user_type = 2", id).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//不存在
			code = types.UserNotExisted
			return
		}
		panic(err.Error())
	}

	if user.Status == 1 {
		//被删除
		code = types.UserHasDeleted
		return
	}
	if err := DB.Raw("select teacher_id,pk_course_id,course_name from course_select_sys_course "+
		"where pk_course_id in(select course_id from course_select_sys_student_select_course "+
		"where student_id = ?)", id).Scan(&course).Error; err != nil {
		panic(err.Error())
	}
	if len(course) == 0 {
		//没课程
		code = types.StudentHasNoCourse
		return
	}
	for _, coursenow := range course {
		tcourse = append(tcourse, types.TCourse{
			CourseID:  strconv.FormatInt(coursenow.PkCourseId, 10),
			Name:      coursenow.CourseName,
			TeacherID: strconv.FormatInt(coursenow.TeacherId, 10),
		})
	}
}
