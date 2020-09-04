package teachers

import (
	"fmt"

)

func handle() {
	fmt.Println("1 ) see test results  \n2 ) login to class \n3 ) sign up to a class ")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:

	case 2:

	case 3:

	default:
		fmt.Println("unsupported input !")
		handle()
	}
}