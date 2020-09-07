package class

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
)

func (class TeacherClass) Handle() {
	fmt.Println("fuckkkkkk")
	//createExam(class)
	//seeResults(class)
	seeRequests(class)
}

//TODO creating exam
func createExam(class TeacherClass) {
}

//TODO showing results of test
func seeResults(class TeacherClass) {
}

//TODO see request of chosen class
func seeRequests(class TeacherClass) {
	reqs := dbHandler.GetRequests(class.Db, class.Id)
	for cl, std := range reqs {
		fmt.Print("REQUEST FROM :  " + std + "  FOR CLASS : " + cl)
		fmt.Println()
	}
	//TODO check class Id to be valid
	//TODO choosing request to be accepted
	var classID string
	fmt.Println("ENTER CLASS ID : ")
	fmt.Scan(&classID)
}
