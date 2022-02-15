package entity

import "time"

type User struct {
	PkUserId   int64  `gorm:"primaryKey,autoIncrement"`
	UkUsername string `gorm:"varchar(32)"`
	//Username   string `gorm:"column:uk_username"`
	Nickname   string `gorm:"varchar(32)"`
	Password   string `gorm:"varchar(64)"`
	UserType   int8
	Status     int8
	UpdateTime time.Time
	CreateTime time.Time
}

func (User) TableName() string {
	return "course_select_sys_user"
}

type Course struct {
	PkCourseId    int64  `gorm:"primaryKey,autoIncrement"`
	CourseName    string `gorm:"varchar(32)"`
	TotalCapacity int
	LeftCapacity  int
	TeacherId     int64
}

func (Course) TableName() string {
	return "course_select_sys_course"
}
