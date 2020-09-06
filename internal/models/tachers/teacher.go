package teachers

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/class"
	"github.com/hawwwdi/Goxam/internal/models/dbHandler"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

func Handle(user user.User, db *sql.DB) {
	fmt.Println("HELLO Mr." + user.UserName + " welcome to your portal .")
	fmt.Println("1 ) login to class \n2 ) create a new class")
	var chosen int
	fmt.Scan(&chosen)
	switch chosen {
	case 1:
		loginToClass(user, db)
	case 2:
		createClass(user, db)
	default:
		fmt.Println("unsupported input !")
		Handle(user, db)
	}
}
func createClass(user user.User, db *sql.DB) {
	var id string
	fmt.Println("ENTER NAME OF THE CLASS ( NOT REPEATED ) : ")
	fmt.Scan(&id)
	id = user.UserName + "_" + id
	dbHandler.SaveClass(db, user, id)
	//TODO going to class
	class.Handle(id)
}
func loginToClass(user user.User, db *sql.DB) {
	classesId := dbHandler.GetTeacherClasses(db, user)
	for _, id := range classesId {
		fmt.Println(" = " + id)
	}
	var id string
	fmt.Println("ENTER NAME OF THE CLASS : ")
	fmt.Scan(&id)
	class.Handle(id)
}
