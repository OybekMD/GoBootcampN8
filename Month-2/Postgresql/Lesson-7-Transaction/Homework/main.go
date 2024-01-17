package main

import (
	m "GoBootcampN8/Month-2/Postgresql/Lesson-7-Transaction/Homework/Models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=ebot dbname=alprov1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Data
	studentDatas := []m.Student{
		{
			Fullname:  "Oybek Atamatov",
			Username:  "MDpro",
		},
		{
			Fullname:  "Doston Shernazarov",
			Username:  "pythonDev",
		},
		{
			Fullname:  "Kozim Jumayev",
			Username:  "sharpist",
		},
	}

	teacherData := []m.Teacher{
		{
			Fullname:  "Nurali Uktamov",
			Username:  "nurali",
		},
		{
			Fullname:  "Axrorbek Olimjonov",
			Username:  "axrorbek",
		},
	}

	courseData := m.Courses{
		Name:      "Golang",
	}

	studentsCoursesData := m.Students_Courses{} // Data add in Insert time !

	studentReqData := m.StudentReq{
		Student:          studentDatas[0],
		Teacher:          teacherData[0],
		Courses:          courseData,
		Students_Courses: studentsCoursesData,
	}
	Create(db, studentReqData)

	// pp.Printf("StudentReq Data:\n%+v\n", studentReqData)
}

func Create(db *sql.DB, student m.StudentReq) error {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	var studentId int
	if err := tx.QueryRow(`INSERT INTO students(fullname, username)
		VALUES ($1, $2) RETURNING id`,
		student.Student.Fullname, student.Student.Username).Scan(&studentId); err != nil {
			tx.Rollback()
			return err
	}
	
	var teacherId int
	if err := tx.QueryRow(`INSERT INTO teachers(fullname, username)
		VALUES ($1, $2) RETURNING id`,
		student.Teacher.Fullname, student.Teacher.Username).Scan(&teacherId); err != nil {
			tx.Rollback()
			return err
	}

	var courseId int
	if err := tx.QueryRow(`INSERT INTO courses(name, teacher)
		VALUES ($1, $2) RETURNING id`,
		student.Courses.Name, teacherId).Scan(&courseId); err != nil {
			tx.Rollback()
			return err
	}

	if _, err := tx.Exec(`INSERT INTO students_courses(student_id, course_id)
		VALUES ($1, $2)`,
		studentId, courseId); err != nil {
			tx.Rollback()
			return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("Sucessfully Created!")
	return nil
}