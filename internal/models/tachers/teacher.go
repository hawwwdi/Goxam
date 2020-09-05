package teachers

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

func Handle(user user.User) {
	fmt.Print("HELLO Mr."+user.Name+" welcome to your portal .")
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) create a new class \n4 ) see students requests ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		seeTestResult(user)
	case 2:
		loginToClass(user)
	case 3:
		createClass(user)
	case 4:
		checkRequests(user)
	default:
		fmt.Println("unsupported input !")
		Handle(user)
	}
}
func createClass(user user.User) {
	//going to class template
	var classId string
	class.Handle(classId )
}
func seeTestResult(user user.User) {
	//show the quizzes that this teacher has been made
	// after choosing one we will show the result of that
}
func loginToClass(user user.User) {
	//going to class template
	var classId string
	class.Handle(classId )
}
func checkRequests(user user.User)  {
	// chek the requests and add the students to class if it was ok
}