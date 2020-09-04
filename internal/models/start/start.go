package start

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/students"
	"github.com/hawwwdi/Goxam/internal/models/tachers"
)


type User struct {
	Name, Email, Password string
	Type                  int//1 teacher , 2 student
}

func start() {
	for true {
		fmt.Println("1 ) sign up \n2 ) login ")
		var chosen int
		fmt.Scan(&chosen)
		switch chosen {
		case 1:
			go signUp()
		case 2:
			go login()
		default:
			print("unsupported input !")
		}
	}
}
func signUp() {
	fmt.Println("signing in ... ")
	newUser := getUser()
	if checkedLog(){
		fmt.Print("You have already logged in ! ")
		login()
	}else {
		saveUserToDB(newUser)
		if newUser.Type==1{
			teachers.Handle(newUser)
		}else {
			students.Handle(newUser)
		}
	}
}
func login() {

	fmt.Println("signing in ... ")
	newUser := getUser()
	if checkedLog(){
		if newUser.Type==1{
			teachers.Handle(newUser)
		}else {
			students.Handle(newUser)
		}
	}else {
		fmt.Println("you have not sign up yet ! pleas sign up first . ")
		signUp()
	}
}
func getUser() User{
	var Name, Email, Password string
	var Type int
	fmt.Print("NAME : ")
	fmt.Scan(&Name)
	fmt.Print("Email : ")
	fmt.Scan(&Email)
	fmt.Print("Password : ")
	fmt.Scan(&Password)
	fmt.Print("Type of user ( 1 stands for teacher and 2 stands for student )  : ")
	fmt.Scan(&Type)
	return User{Name: Name, Email: Email, Password: Password, Type: Type}
}
func saveUserToDB(user User)  {
	//save user to data base
}
func checkedLog() bool  {
	//search in database and return true if user was logged in before
	return true
}
