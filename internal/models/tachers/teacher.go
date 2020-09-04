package teachers

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/start"
	"github.com/hawwwdi/Goxam/internal/models/class"
)

func Handle(user start.User) {
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
func createClass(user start.User) {
	//going to class template
	var classId string
	class.Handle(user  , classId )
}
func seeTestResult(user start.User) {
	//show the quizzes that this teacher has been made
	// after choosing one we will show the result of that
}
func loginToClass(user start.User) {
	//going to class template
	var classId string
	class.Handle(user  , classId )
}
func checkRequests(user start.User)  {
	// chek the requests and add the students to class if it was ok
}