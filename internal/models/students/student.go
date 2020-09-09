package students

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/db/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/user"
)



type Student struct {
	UserName, Email string
	PassWord        []byte
}


func Handle(user user.User) {
	fmt.Println("Hello student " + user.UserName + " Welcome to your portal . ")
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) sign up to a class ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		seeTestResult(user)
	case 2:
		loginToClass(user)
	case 3:
		signUpToClass(user)
	default:
		fmt.Println("unsupported input !")
		Handle(user)
	}
}
func signUpToClass(user user.User) {
	var classId string
	fmt.Scan(&classId)
	massage := dbHandler.CheckClass(user, classId)
	if massage == "ok" {
		dbHandler.SendRequest( user, classId)
	} else {
		fmt.Println(massage)
	}
	Handle(user)
}
func seeTestResult(user user.User) {
	//TODO
	//show the test that student has participated in
	//after choosing one we will sho the result of that
}
func loginToClass(user user.User) {
	classesId := dbHandler.GetStudentsClasses(user)
	for _, id := range classesId {
		fmt.Println(" = " + id)
	}
	var id string
	fmt.Println("ENTER NAME OF THE CLASS : ")
	fmt.Scan(&id)
	_,ok:= classesId[id]
	if ok{
		class.StudentClass{Id: id, User: user}.Handle()
	}else {
		fmt.Println("unknown id")
		loginToClass(user)
	}
}
