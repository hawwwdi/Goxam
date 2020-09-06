package dbHandler

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/user"
	"log"
)
// getting data base from connection and return it to app
func GetDB() *sql.DB {
	db, err := sql.Open("mysql",
		"root:armin3011@tcp(127.0.0.1:3306)/Goxam")
	errHandler(err)
	err = db.Ping()
	errHandler(err)
	return db
}

//#################################################################       SAVING USERS ON DATA BASE :
func SaveTeacher(db *sql.DB, user user.User) {
	stmt, err := db.Prepare(`INSERT INTO teachers VALUES (?,?,?);`)
	errHandler(err)
	saveUser(stmt, user)
}

func SaveStudent(db *sql.DB, user user.User) {
	stmt, err := db.Prepare(`INSERT INTO students VALUES (?,?,?);`)
	errHandler(err)
	saveUser(stmt, user)
}

func saveUser(stmt *sql.Stmt, user user.User) {
	defer stmt.Close()
	r, err := stmt.Exec(user.Email, user.UserName, string(user.PassWord))
	errHandler(err)
	ro, err := r.RowsAffected()
	errHandler(err)
	fmt.Println("INSERTED RECORD to students or teachers", ro)
}
//################################################################################################# END PART /

//################################################################## LOADING USERS FROM DATA BASE :
func LoadUsersFromDb(db *sql.DB, teachers map[string]user.User, students map[string]user.User) {
	rows, err2 := db.Query("SELECT * FROM Goxam.teachers")
	errHandler(err2)
	putUsers(rows, teachers, 1)
	rows2, err2 := db.Query("SELECT * FROM Goxam.students")
	errHandler(err2)
	putUsers(rows2, students, 2)
}
func putUsers(rows *sql.Rows, users map[string]user.User, typee int) {
	var e, n, p string
	for rows.Next() {
		err := rows.Scan(&e, &n, &p)
		errHandler(err)
		newUser := user.User{Email: e, UserName: n, PassWord: []byte(p), Type: typee}
		users[newUser.Email] = newUser
	}
}
//################################################################################################## END PART /
func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}