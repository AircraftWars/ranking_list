package auth

type User struct {
	PkUserId   int64  `gorm:"column:user_id"`
	UkUsername string `gorm:"column:user_name"`
	Nickname   string
	Password   string `gorm:"column:password"`
	UserType   int8   `gorm:"column:user_type"`
	Status     int8
}

func (User) TableName() string {
	return "users"
}

type Course struct {
	PkCourseId    int64  `gorm:"primaryKey"`
	CourseName    string `gorm:"varchar(32)"`
	TotalCapacity int
	LeftCapacity  int
	TeacherId     int64
}

func (Course) TableName() string {
	return "course_select_sys_course"
}
