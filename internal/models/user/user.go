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


