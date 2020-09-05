package start

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/students"
	"github.com/hawwwdi/Goxam/internal/models/tachers"
	"github.com/hawwwdi/Goxam/internal/models/user"
	"golang.org/x/crypto/bcrypt"
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
	newUser := user.User{Name: Name, Email: Email, PassWord: []byte(Password), Type: Type}
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
	user.SetEncryptPassWord(user.PassWord)
	if user.Type == 1 {
		Teachers[user.Email] = user
		dbHandler.SaveTeacher(db, user)
	} else {
		Students[user.Email] = user
		dbHandler.SaveStudent(db, user)
	}
}

//#########################################################################################################
func login() {
	fmt.Println("logging in ... ")
	newUser := getUser()
	if checkSignUp(newUser) {
		if checkPass(newUser) {
			if newUser.Type == 1 {
				teachers.Handle(newUser)
			} else {
				students.Handle(newUser)
			}
		} else {
			fmt.Print("WRONG PASS WORD")
		}
	} else {
		fmt.Println("you have not sign up yet ! pleas sign up first . ")
		signUp()
	}
}

func checkPass(user user.User) bool {
	err := bcrypt.CompareHashAndPassword(Students[user.Email].PassWord, user.PassWord)
	if err != nil {
		return false
	}
	return true
}
//#########################################################################################################
func checkSignUp(user user.User) bool {
	var ok bool
	if user.Type == 1 {
		_, ok = Teachers[user.Email]
	} else {
		_, ok = Students[user.Email]
	}
	if !ok {
		return false
	} else {
		return true
	}
}

//##########################################################################################################

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
