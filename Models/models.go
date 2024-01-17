package models

import "time"

type Student struct {
	Fullname  string
	Username  string
	CreatedAt time.Time
}

type Teacher struct {
	Fullname  string
	Username  string
	CreatedAt time.Time
}

type Courses struct {
	Name      string
	Teacher   int
	CreatedAt time.Time
}

type Students_Courses struct {
	Student_id int
	Course_id  int
}

type StudentReq struct {
	Student Student
	Teacher Teacher
	Courses Courses
	Students_Courses Students_Courses 
}
