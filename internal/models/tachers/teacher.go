package teachers

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/db/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)



type Teacher struct {
	UserName, Email string
	PassWord        []byte
}
// function to create encrypted password
func (teacher *Teacher) SetEncryptPassWord(pass []byte) {
	bs, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	teacher.PassWord = bs
}

func (teacher *Teacher)Handle() {
	fmt.Println("HELLO Mr." + teacher.UserName + " welcome to your portal .")
	fmt.Println("1 ) login to class \n2 ) create a new class")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		teacher.loginToClass()
	case 2:
		teacher.createClass()
	default:
		fmt.Println("unsupported input !")
		teacher.Handle()
	}
}
func (teacher *Teacher)createClass() {
	var id string
	fmt.Println("ENTER NAME OF THE CLASS ( NOT REPEATED ) : ")
	fmt.Scan(&id)
	id = teacher.UserName + "_" + id
	dbHandler.SaveClass( user, id)
	class.TeacherClass{Id: id, User: user}.Handle()
}
func (teacher *Teacher)loginToClass() {
	classesId := dbHandler.GetTeacherClasses(user)
	for _, id := range classesId {
		fmt.Println(" = " + id)
	}
	var id string
	fmt.Println("ENTER NAME OF THE CLASS : ")
	fmt.Scan(&id)
	_,ok:=classesId[id]
	if ok {
		class.TeacherClass{Id: id, User: user}.Handle()
	}else {
		fmt.Println("unknown id")
		teacher.loginToClass()
	}
}
