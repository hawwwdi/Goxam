package students

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

func Handle(user user.User, db *sql.DB) {
	fmt.Println("Hello student " + user.UserName + " to your portal . ")
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) sign up to a class ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		seeTestResult(user)
	case 2:
		loginToClass(user, db)
	case 3:
		signUpToClass(user, db)
	default:
		fmt.Println("unsupported input !")
		Handle(user, db)
	}
}
func signUpToClass(user user.User, db *sql.DB) {
	var classId string
	fmt.Scan(&classId)
	massage := dbHandler.CheckClass(db,user,classId)
	if massage=="ok"{
		dbHandler.SendRequest(db,user,classId)
	} else {
		fmt.Println(massage)
	}
	Handle(user, db)
}
func seeTestResult(user user.User) {
	//TODO
	//show the test that student has participated in
	//after choosing one we will sho the result of that
}
func loginToClass(user user.User, db *sql.DB) {
	classesId := dbHandler.GetStudentsClasses(db, user)
	for _, id := range classesId {
		fmt.Println(" = " + id)
	}
	var id string
	fmt.Println("ENTER NAME OF THE CLASS : ")
	fmt.Scan(&id)
	class.Handle(id)
}
