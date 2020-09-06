package students

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

func Handle(user user.User) {
	fmt.Println("Hello student " + user.UserName + " to your portal . ")
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
	//get name of teacher
	//get name or id of class
	//save this user to list of requests of the teacher
	Handle(user)
}
func seeTestResult(user  user.User) {
	//show the test that student has participated in
	//after choosing one we will sho the result of that
}
func loginToClass(user  user.User) {
	//show the classes that student is in
	//go to class template
	var classId string
	class.Handle(classId)
}
