package main

import (
	. "GoBootcampN8/Month-2/Postgresql/Lesson-6-GORM/Homework/Models"
	

	pp "github.com/k0kubun/pp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func autoMigrate() {
	db.AutoMigrate(&Students{})
	db.AutoMigrate(&Languages{})
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
	db.Preload("Languages").First(&student, studentID)
	return student
}

func GetAllStudents() []Students {
	var students []Students
	db.Preload("Languages").Find(&students)
	return students
}
func GetAllLanguages() []Languages {
	var languages []Languages
	db.Preload("Languages").Find(&languages)
	return languages
}
func createCourse(name string, price int) {
	course := Languages{
		Name:  name,
		Price: price,
	}
	db.Create(&course)
}
func StudenJoinToCourse(studentID, courseID uint) {
	var student Students
	var course Languages
	db.First(&student, studentID)
	db.First(&course, courseID)
	db.Model(&student).Association("Languages").Append(&course)
}
func main() {
	dsn := "user=postgres password=ebot dbname=gormlesson sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigrate()
	// createStudent("Oybek Atamatov", 20)
	// createStudent("Doston Shernazaor", 22)
	// createStudent("Davron Nuriddinov", 20)
	// createStudent("Ali Abdullayev", 20)
	// createCourse("English", 880000)
	// createCourse("Russian", 65000)
	// StudenJoinToCourse(1, 1)
	// StudenJoinToCourse(2, 1)
	// StudenJoinToCourse(2, 2)
	// StudenJoinToCourse(3, 1)
	// StudenJoinToCourse(3, 2)
	// StudenJoinToCourse(4, 1)
	// updateStudent(1, "Umit Hakimit", 99)
	// deleteStudent (4)
	// student := getStudentByID(1)
	// pp.Println(student)
	allStudents := GetAllStudents()
	pp.Println(allStudents)
	languages := GetAllLanguages()
	pp.Println(languages)
}
