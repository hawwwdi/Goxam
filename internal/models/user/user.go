package user

type User interface {
	SetEncryptPassWord(pass []byte)
	Handle()
	signUpToClass()
	seeTestResult()
	loginToClass()
}
