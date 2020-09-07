package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName, Email string
	PassWord        []byte
	Type            int //1 teacher , 2 student
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

