package start

import (
	"fmt"

	"github.com/hawwwdi/Goxam/internal/db/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/students"
	"github.com/hawwwdi/Goxam/internal/models/tachers"
	"github.com/hawwwdi/Goxam/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

var Teachers = make(map[string]user.User)
var Students = make(map[string]user.User)

// loading users from data base and starting the app
func Start() {
	dbHandler.LoadUsersFromDb(Teachers, Students)
	fmt.Println(Teachers)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println(Students)
	for true {
		fmt.Println("1 ) sign up \n2 ) login ")
		var chosen int
		fmt.Scan(&chosen)
		switch chosen {
		//TODO
		// adding go before each func cause capabality to handle lots of requests
		case 1:
			signUp()
		case 2:
			login()
		default:
			print("unsupported input !")
		}
	}
	defer dbHandler.CloseDB()
}

//##########################################################################################################
// a console version for getting user info
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
	newUser := user.User{UserName: Name, Email: Email, PassWord: []byte(Password), Type: Type}
	return newUser
}

//##########################################################################################################
// signup part
func signUp() {
	fmt.Println("signing in ... ")
	newUser := getUser()
	if checkSignUp(newUser) {
		fmt.Print("You have already signed in ! ")
		login()
	} else {
		saveUserToDB(newUser)
		if newUser.Type == 1 {
			teachers.Handle(newUser)
		} else {
			students.Handle(newUser)
		}
	}
}

// function to save new user to data base and our maps of users
func saveUserToDB(user user.User) {
	user.SetEncryptPassWord(user.PassWord)
	if user.Type == 1 {
		Teachers[user.Email] = user
		dbHandler.SaveTeacher(user)
	} else {
		Students[user.Email] = user
		dbHandler.SaveStudent(user)
	}
}

//#########################################################################################################
//method for logging in
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
			fmt.Println("WRONG PASS WORD")
		}
	} else {
		fmt.Println("you have not sign up yet ! pleas sign up first . ")
		signUp()
	}
}

//this function check inputted password to be correct
func checkPass(user user.User) bool {
	var err error
	if user.Type == 1 {
		err = bcrypt.CompareHashAndPassword(Teachers[user.Email].PassWord, user.PassWord)
	} else {
		err = bcrypt.CompareHashAndPassword(Students[user.Email].PassWord, user.PassWord)
	}
	if err != nil {
		return false
	}
	return true
}

//#########################################################################################################
//this method check if the user exists in our data base or not
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
