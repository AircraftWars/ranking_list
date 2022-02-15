package auth

type User struct {
	PkUserId   int64  `gorm:"primaryKey"`
	UkUsername string `gorm:"varchar(20)"`
	Nickname   string `gorm:"varchar(20)"`
	Password   string `gorm:"varchar(20)"`
	UserType   int8
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
