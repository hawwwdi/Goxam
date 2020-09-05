package start

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/students"
	"github.com/hawwwdi/Goxam/internal/models/tachers"
	"github.com/hawwwdi/Goxam/internal/models/user"
	"log"
)

var db *sql.DB
var Teachers = make(map[string]user.User)
var Students = make(map[string]user.User)

func Start() {
	db = dbHandler.GetDB()
	dbHandler.LoadUsersFromDb(db, Teachers, Students)
	fmt.Println(Teachers)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println(Students)
	for true {
		fmt.Println("1 ) sign up \n2 ) login ")
		var chosen int
		fmt.Scan(&chosen)
		switch chosen {
		//TODO
		// adding go before each func cause capabality to handle lots of reauests
		case 1:
			signUp()
		case 2:
			login()
		default:
			print("unsupported input !")
		}
	}
	defer db.Close()
}
//##########################################################################################################
func getUser() user.User {
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
	newUser := user.User{Name: Name, Email: Email, PassWord: nil, Type: Type}
	newUser.SetEncryptPassWord(Password)
	return newUser
}
//##########################################################################################################
// signup part done
func signUp() {
	fmt.Println("signing in ... ")
	newUser := getUser()
	if checkSignUp(newUser) {
		fmt.Print("You have already signed in ! ")
		login()
	} else {
		saveUserToDB(db, newUser)
		if newUser.Type == 1 {
			teachers.Handle(newUser)
		} else {
			students.Handle(newUser)
		}
	}
}

func saveUserToDB(db *sql.DB, user user.User) {
	if user.Type == 1 {
		dbHandler.SaveTeacher(db, user)
	} else {
		dbHandler.SaveStudent(db, user)
	}
}
//#########################################################################################################
func login() {
	fmt.Println("logging in ... ")
	newUser := getUser()
	if checkSignUp(newUser) {
		if newUser.Type == 1 {
			teachers.Handle(newUser)
		} else {
			students.Handle(newUser)
		}
	} else {
		fmt.Println("you have not sign up yet ! pleas sign up first . ")
		signUp()
	}
}
//#########################################################################################################
func checkSignUp(user user.User) bool {
	var ok  bool
	if user.Type==1{
		_ ,ok = Teachers[user.Email]
	}else {
		_, ok = Students[user.Email]
	}
	if !ok{
		return false
	}else {
		return true
	}
}
//##########################################################################################################

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
