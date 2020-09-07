package class

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
)

func (class TeacherClass) Handle() {
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
	for std, cl := range reqs {
		fmt.Print(" REQUEST FROM :  " + std + "  FOR CLASS : " + cl)
		fmt.Println()
	}
	//TODO choosing multiple request to be accepted
	var email string
	fmt.Println("ENTER STUDENT EMAIL : ")
	fmt.Scan(&email)
	//TODO check student email to be valid
	dbHandler.AddStudentTOCLass(class.Id,email,class.Db)
}
