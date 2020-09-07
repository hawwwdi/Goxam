package class

import (
	"database/sql"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

type TeacherClass struct {
	Id string
	User user.User
}
type StudentClass struct {
	Id string
	User user.User
	Db *sql.DB
}
type HandleClass interface {
	Handle()
}

