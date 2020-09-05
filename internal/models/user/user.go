package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Name, Email string
	PassWord    []byte
	Type        int //1 teacher , 2 student
}
// function to create encripted password
func (user *User) SetEncryptPassWord(pass []byte) {
	bs, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
		return
	}
	user.PassWord = bs
}

