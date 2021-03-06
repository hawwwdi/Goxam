package dbHandler

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/user"
)


//////////////////////////////////////////////////////////////////////// TEACHERS PART :
//save new classes to db
func SaveClass(user user.User, id string) {
	stmt, err := db.Prepare(`INSERT INTO classes VALUES (?,?);`)
	errHandler(err)
	defer stmt.Close()
	r, err1 := stmt.Exec(id, user.Email)
	errHandler(err1)
	ro, err2 := r.RowsAffected()
	errHandler(err2)
	fmt.Println("INSERTED RECORD TO Classes", ro)
}

//get each teacher classes from db
func GetTeacherClasses( user user.User) map[string]string {
	rows, err2 := db.Query("SELECT class_id FROM Goxam.classes WHERE teacher_email= ?", user.Email)
	errHandler(err2)
	return getClasses(rows)
}

//get a class and return list of requests to join to class
func GetRequests( id string) map[string]string {
	var reqs = make(map[string]string)
	rows, err2 := db.Query("SELECT student_email FROM Goxam.requests WHERE class_id= ?", id)
	errHandler(err2)
	for rows.Next() {
		var std string
		err := rows.Scan(&std)
		reqs[std] = id
		errHandler(err)
	}
	return reqs
}
func AddStudentTOCLass(id, email string) {
	stmt, err := db.Prepare(`INSERT INTO class_participation VALUES (?,?);`)
	errHandler(err)
	defer stmt.Close()
	r, err1 := stmt.Exec(id, email)
	errHandler(err1)
	ro, err2 := r.RowsAffected()
	errHandler(err2)
	fmt.Println("INSERTED RECORD to classParticipation", ro)
	//removing request after accepting it
	RemoveRequest(id,email)
}
func RemoveRequest(id , email string)  {
	res, err := db.Exec("DELETE FROM requests WHERE class_id=? AND  student_email = ?", id , email)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			fmt.Println(count)
		}
	}
}
////////////////////////////////////////////////////////////////////////// END OF PART /
//////////////////////////////////////////////////////////////////////// STUDENTS PART :
func GetStudentsClasses( user user.User) map[string]string {
	rows, err2 := db.Query("SELECT class_id FROM Goxam.class_participation WHERE student_email= ?", user.Email)
	errHandler(err2)
	return getClasses(rows)
}

//check if the wanted class exists and check if student is already in (using checkParticipated)
func CheckClass( user user.User, id string) string {
	rows, err2 := db.Query("SELECT EXISTS(SELECT class_id FROM Goxam.classes WHERE class_id= ?)", id)
	errHandler(err2)
	var existence string
	for rows.Next() {
		err := rows.Scan(&existence)
		errHandler(err)
	}
	if existence == "0" {
		return "could not find the class ! "
	} else if checkParticipated(user, id, db) {
		return "your already in class ! "
	} else {
		return "ok"
	}
}
//check if student is already in class
func checkParticipated(user user.User, id string, db *sql.DB) bool {
	classes := GetStudentsClasses( user)
	for _, value := range classes {
		if value == id {
			return true
		}
	}
	return false
}

//after checkClass , save class id and student email on requests table
func SendRequest(user user.User, id string) {
	stmt, err := db.Prepare(`INSERT INTO requests VALUES (?,?);`)
	errHandler(err)
	defer stmt.Close()
	r, err1 := stmt.Exec(user.Email, id)
	errHandler(err1)
	ro, err2 := r.RowsAffected()
	errHandler(err2)
	fmt.Println("INSERTED RECORD to requests", ro)
}

////////////////////////////////////////////////////////////////////////// END OF PART /
func getClasses(rows *sql.Rows) map[string]string {
	var classesId = make(map[string]string)
	var classId string
	for rows.Next() {
		err := rows.Scan(&classId)
		classesId[classId] = classId
		errHandler(err)
	}
	return classesId
}
