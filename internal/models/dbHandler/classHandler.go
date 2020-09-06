package dbHandler

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/user"
)
//////////////////////////////////////////////////////////////////////// TEACHERS PART :
//save new classes to db
func SaveClass(db *sql.DB, user user.User, id string) {
	stmt, err := db.Prepare(`INSERT INTO classes VALUES (?,?);`)
	errHandler(err)
	defer stmt.Close()
	r, err := stmt.Exec(id, user.Email)
	errHandler(err)
	ro, err := r.RowsAffected()
	errHandler(err)
	fmt.Println("INSERTED RECORD TO Classes", ro)
}
//get each teacher classes from db
func GetTeacherClasses(db *sql.DB, user user.User) map[int]string {
	var classesId = make(map[int]string)
	rows, err2 := db.Query("SELECT class_id FROM Goxam.classes WHERE teacher_email= ?", user.Email)
	errHandler(err2)
	var classId string
	var counter = 0
	for rows.Next() {
		err := rows.Scan(&classId)
		classesId[counter] = classId
		counter += 1
		errHandler(err)
	}
	return classesId
}
////////////////////////////////////////////////////////////////////////// END OF PART /