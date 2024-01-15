package main

import (
	"database/sql"
	"fmt"

	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
)

type Student struct {
	ID        int
	Name      string
	Birthyear int
}

type Group struct {
	ID      int
	Name    string
	Teacher string
}

type Student_Courses struct {
	ID      int
	Student int
	Group   int
}

func main() {
	students := []Student{
		Student{Name: "Oybek Atamatov", Birthyear: 2003},
		Student{Name: "Doston Shernazarov", Birthyear: 2001},
		Student{Name: "Nodirbek Numanov", Birthyear: 2003},
		Student{Name: "Ulugbek Orazimbekov", Birthyear: 2002},
		Student{Name: "Xumoyinmurzo Qoshmaqboyev", Birthyear: 2003},
	}
	groups := []Group{
		Group{Name: "Bootcamp Foundation", Teacher: "Islom Abdurahmonov"},
		Group{Name: "GO bootcamp N8", Teacher: "Nurali Uktamov"},
	}

	// Show our all print dates
	pp.Println(students)
	pp.Println(groups)

	fmt.Println("<<<----- Function Results ----->>>")

	connStr := "user=postgres password=ebot dbname=najottalimhomew sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Insert [ ---- ]
	// CreateStudent(db, &students[0]) // from 0 to 4
	// CreateStudent(db, &students[1]) // from 0 to 4
	// CreateStudent(db, &students[2]) // from 0 to 4
	// CreateStudent(db, &students[3]) // from 0 to 4
	// CreateStudent(db, &students[4]) // from 0 to 4

	// CreateGroup(db, &groups[0]) // 0 and 1
	// CreateGroup(db, &groups[1]) // 0 and 1

	// StudentJoinToGroup(db, 1, 1)
	// StudentJoinToGroup(db, 2, 1)
	// StudentJoinToGroup(db, 4, 1)
	// StudentJoinToGroup(db, 5, 1)

	// StudentJoinToGroup(db, 1, 2)
	// StudentJoinToGroup(db, 2, 2)
	// StudentJoinToGroup(db, 3, 2)
	// StudentJoinToGroup(db, 4, 2)
	// StudentJoinToGroup(db, 5, 2)

	// Updates [ ---- ]
	// UpdateStudent(db, &Student{Name: "Eshmatov Toshmat", Birthyear: 2000}, 2)
	// UpdateGroup(db, &Group{Name: "Go Bootcamp updated", Teacher: "Axrorbek Olimjonov"}, 1)

	// Select All [ ---- ]
	// GetStudent(db, 1)
	// GetGroup(db, 1)
	// GetAllStudent(db)
	// GetAllGroup(db)
	// GetStudentWithGroup(db, 2)
	// GetStudentsOfGroup(db, 1)

	// Delete [ ---- ]
	// DeleteGroup(db, 2)
	// DeleteStudent(db, 3)
	// StudentRemoveFromGroup(db, 2)
}

func CreateStudent(db *sql.DB, student *Student) error {
	var respStudent Student
	rowStudent := db.QueryRow(`INSERT INTO students (name, birthyear) VALUES ($1, $2) returning id`, student.Name, student.Birthyear)
	if err := rowStudent.Scan(&respStudent.ID); err != nil {
		return fmt.Errorf("failed to create group: %v", err)
	}
	fmt.Println("Student sucessfuly inserted with id:", respStudent.ID)
	return nil
}

func CreateGroup(db *sql.DB, group *Group) error {
	var respGroup Group
	rowGroup := db.QueryRow(`INSERT INTO Groups (name, teacher) VALUES ($1, $2) returning id`, group.Name, group.Teacher)
	if err := rowGroup.Scan(&respGroup.ID); err != nil {
		return fmt.Errorf("failed to create group: %v", err)
	}
	fmt.Println("Group sucessfuly inserted with id:", respGroup.ID)
	return nil
}

func UpdateStudent(db *sql.DB, student *Student, id int) error {
	result, err := db.Exec("update students set name = $1, birthyear = $2  where id = $3", student.Name, student.Birthyear, id)
	if err != nil {
		return fmt.Errorf("failed to update student with ID %d: %v", id, err)
	}
	fmt.Print("Student Sucessfuly Updated: ")
	fmt.Println(result.RowsAffected())
	return nil
}

func UpdateGroup(db *sql.DB, group *Group, id int) error {
	result, err := db.Exec("update groups set name = $1, teacher = $2  where id = $3", group.Name, group.Teacher, id)
	if err != nil {
		return fmt.Errorf("failed to update student with ID %d: %v", id, err)
	}
	fmt.Print("Group Sucessfuly Updated: ")
	fmt.Println(result.RowsAffected())
	return nil
}

func GetStudent(db *sql.DB, id int) error {
	var column1Value, column2Value string
	err := db.QueryRow("SELECT name, birthyear FROM students WHERE id = $1", id).Scan(&column1Value, &column2Value)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case err != nil:
		return fmt.Errorf("no student found with ID %d", id)
	default:
		fmt.Printf("name: %s, birthyear: %s\n", column1Value, column2Value)
	}
	return nil
}
func GetGroup(db *sql.DB, id int) error {
	var column1Value, column2Value string
	err := db.QueryRow("SELECT name, teacher FROM groups WHERE id = $1", id).Scan(&column1Value, &column2Value)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case err != nil:
		return fmt.Errorf("no group found with ID %d", id)
	default:
		fmt.Printf("name: %s, teacher: %s\n", column1Value, column2Value)
	}
	return nil
}

func GetAllStudent(db *sql.DB) error {
	rows, err := db.Query("select * from students")
	if err != nil {
		return fmt.Errorf("failed to fetch stundets: %v", err)
	}
	defer rows.Close()
	students := []Student{}

	for rows.Next() {
		student := Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Birthyear)
		if err != nil {
			fmt.Println(err)
			continue
		}
		students = append(students, student)
	}
	for _, s := range students {
		fmt.Println(s.ID, s.Name, s.Birthyear)
	}
	return nil
}

func GetAllGroup(db *sql.DB) error {
	rows, err := db.Query("select * from groups")
	if err != nil {
		return fmt.Errorf("failed to fetch groups: %v", err)
	}
	defer rows.Close()
	groups := []Group{}

	for rows.Next() {
		group := Group{}
		err := rows.Scan(&group.ID, &group.Name, &group.Teacher)
		if err != nil {
			fmt.Println(err)
			continue
		}
		groups = append(groups, group)
	}
	for _, g := range groups {
		fmt.Println(g.ID, g.Name, g.Teacher)
	}
	return nil
}

func DeleteStudent(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM student_group WHERE student_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete references from student_group for student with ID %d: %v", id, err)
	}

	_, err = db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete students with ID %d: %v", id, err)
	}

	fmt.Printf("Student and associated references successfully deleted with ID: %d\n", id)
	return nil
}

func DeleteGroup(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM student_group WHERE group_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete references from student_group for group with ID %d: %v", id, err)
	}

	_, err = db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete group with ID %d: %v", id, err)
	}

	fmt.Printf("Group and associated references successfully deleted with ID: %d\n", id)
	return nil
}

func StudentJoinToGroup(db *sql.DB, student_id int, group_id int) error {
	var respStudentGroup Student_Courses
	rowStudent := db.QueryRow(`INSERT INTO student_group (student_id, group_id) VALUES ($1, $2) returning id`, student_id, group_id)
	if err := rowStudent.Scan(&respStudentGroup.ID); err != nil {
		return err
	}
	fmt.Println("Student sucessfuly inserted to Group with id:", respStudentGroup.ID)
	return nil
}

func StudentRemoveFromGroup(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM student_group WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete students with ID %d: %v", id, err)
	}

	fmt.Printf("Student successfully deleted from group with ID: %d\n", id)
	return nil
}

func GetStudentWithGroup(db *sql.DB, id int) error {
	var column1Value, column2Value, column3Value, column4Value string
	err := db.QueryRow("SELECT s.id, s.name, g.name, g.teacher FROM student_group sg JOIN students s ON sg.id = s.id JOIN groups g ON sg.id = g.id WHERE s.id = $1", id).Scan(&column1Value, &column2Value, &column3Value, &column4Value)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case err != nil:
		return fmt.Errorf("no student found with ID %d", id)
	default:
		fmt.Printf("name: %s, birthyear: %s, groupname: %s, teacher: %s \n", column1Value, column2Value, column3Value, column4Value)
	}
	return nil	
}

func GetStudentsOfGroup(db *sql.DB, id int) ([]Student, error) {
	rows, err := db.Query(`
		SELECT s.id, s.name, s.birthyear, g.name, g.teacher
		FROM student_group sg
		JOIN students s ON s.id = sg.student_id 
		JOIN groups g ON g.id = sg.group_id
		WHERE sg.group_id = $1
	`, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get students of group: %v", err)
	}
	defer rows.Close()

	students := make([]Student, 0)
	groups := make([]Group, 0)
	for rows.Next() {
		var student Student
		var group Group
		if err := rows.Scan(&student.ID, &student.Name, &student.Birthyear, &group.Name, &group.Teacher); err != nil {
			return nil, fmt.Errorf("failed to scan student row: %v", err)
		}
		students = append(students, student)
		groups = append(groups, group)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over student rows: %v", err)
	}
	for _, g := range groups {
		fmt.Println(g.ID, g.Name, g.Teacher)
		break
	}
	for _, s := range students {
		fmt.Println(s.ID, s.Name, s.Birthyear)
	}
	return students, nil
}
