package main

import (
	pp "github.com/k0kubun/pp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Students struct {
	Name    string
	Age     int
	Courses []Courses `gorm:"many2many:courses_students"`
	gorm.Model
}
type Courses struct {
	Name     string
	Price    int
	Students []Students `gorm:"many2many:courses_students"`
	gorm.Model
}

func autoMigrate() {
	db.AutoMigrate(&Students{})

	db.AutoMigrate(&Courses{})
	db.AutoMigrate(&StudentCourse{})
}

type StudentCourse struct {
	StudentID uint
	CourseID  uint
	gorm.Model
}

func createStudent(name string, age int) {
	student := Students{
		Name: name,
		Age:  age,
	}
	db.Create(&student)
}

func updateStudent(studentID uint, name string, age int) {
	var student Students
	db.First(&student, studentID)
	student.Name = name
	student.Age = age
	db.Save(&student)
}

func deleteStudent(studentID uint) {
	var student Students
	db.Delete(&student, studentID)
}

func getStudentByID(studentID uint) Students {
	var student Students
	db.Preload("Courses").First(&student, studentID)
	return student
}

func getAllStudents() []Students {
	var students []Students
	db.Preload("Courses").Find(&students)
	return students
}
func getAllCourses() []Courses {
	var courses []Courses
	db.Preload("Courses").Find(&courses)
	return courses
}
func createCourse(name string, price int) {
	course := Courses{
		Name:  name,
		Price: price,
	}
	db.Create(&course)
}
func enrollStudentInCourse(studentID, courseID uint) {
	var student Students
	var course Courses
	db.First(&student, studentID)
	db.First(&course, courseID)
	db.Model(&student).Association("Courses").Append(&course)
}
func main() {
	dsn := "host=localhost_user=postgres password=ebot dbname=gormlesson sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// autoMigrate()
	// createStudent("Sardordek Keldimurodov", 24)
	// createStudent("Xasan Nosirov", 17)
	// createStudent("Davron Nuriddinov", 20)
	// createCourse(".Net", 6200000)
	// createCourse("Java", 7000000) // enrollStudentInCourse(3, 1) // enrollStudent InCourse(1, 2) // enrollStudent InCourse(2, 1) // enrollStudentInCourse(2, 2)
	// updateStudent(1, "Awezone", 26) // deleteStudent (3)
	// student := getStudentByID(1)
	// pp.Println(student)
	allStudents := getAllStudents()
	pp.Println(allStudents)
	courses := getAllCourses()
	pp.Println(courses)
}
