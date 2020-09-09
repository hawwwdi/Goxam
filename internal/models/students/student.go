package students

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/db/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	UserName, Email string
	PassWord        []byte
}
// function to create encrypted password
func (std *Student) SetEncryptPassWord(pass []byte) {
	bs, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	std.PassWord = bs
}

func (std *Student) Handle() {
	fmt.Println("Hello student " + std.UserName + " Welcome to your portal . ")
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) sign up to a class ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		std.seeTestResult()
	case 2:
		std.loginToClass()
	case 3:
		std.signUpToClass()
	default:
		fmt.Println("unsupported input !")
		std.Handle()
	}
}
func (std *Student) signUpToClass() {
	var classId string
	fmt.Scan(&classId)
	massage := dbHandler.CheckClass(user, classId)
	if massage == "ok" {
		dbHandler.SendRequest(user, classId)
	} else {
		fmt.Println(massage)
	}
	Handle(user)
}
func (std *Student) seeTestResult() {
	//TODO
	//show the test that student has participated in
	//after choosing one we will sho the result of that
}
func (std *Student) loginToClass() {
	classesId := dbHandler.GetStudentsClasses(std)
	for _, id := range classesId {
		fmt.Println(" = " + id)
	}
	var id string
	fmt.Println("ENTER NAME OF THE CLASS : ")
	fmt.Scan(&id)
	_, ok := classesId[id]
	if ok {
		class.StudentClass{Id: id, User: user}.Handle()
	} else {
		fmt.Println("unknown id")
		std.loginToClass()
	}
}
