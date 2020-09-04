package students

import "fmt"

func handle() {
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) sign up to a class ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		seeTestResult()
	case 2:
		loginToClass()
	case 3:
		signUpToClass()
	default:
		fmt.Println("unsupported input !")
		handle()
	}
}
func signUpToClass() {
}
func seeTestResult()  {
}
func loginToClass()  {
}
