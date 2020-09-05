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

func (user *User) SetEncryptPassWord(pass string) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
		return
	}
	user.PassWord = bs
}
func (user *User) CheckPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword(user.PassWord, []byte(pass))
	if err != nil {
		return false
	}
	return true
}
