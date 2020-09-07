package class

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/db/dbHandler"
)

func (class TeacherClass) Handle() {
	//createExam(class)
	//seeResults(class)
	checkRequests(class)
}

//TODO creating exam
func createExam(class TeacherClass) {
}

//TODO showing results of test
func seeResults(class TeacherClass) {
}

//see requests for a class and add students to class based on them
func checkRequests(class TeacherClass) {
	reqs := dbHandler.GetRequests(class.Id)
	for std, cl := range reqs {
		fmt.Print(" REQUEST FROM :  " + std + "  FOR CLASS : " + cl)
		fmt.Println()
	}
	//TODO choosing multiple request to be accepted
	var email string
	fmt.Println("ENTER STUDENT EMAIL : ")
	fmt.Scan(&email)
	_,ok := reqs[email]
	if ok{
		dbHandler.AddStudentTOCLass(class.Id,email)
	}
}
