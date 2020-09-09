package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os/user"
)

type User interface {
	SetEncryptPassWord(pass []byte)
	Handle()
	signUpToClass()
	seeTestResult()
	loginToClass()
}
// function to create encrypted password
func (user *User) SetEncryptPassWord(pass []byte) {
	bs, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.PassWord = bs
}

