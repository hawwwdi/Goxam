package class

import (
	"github.com/hawwwdi/Goxam/internal/models/user"
)

type TeacherClass struct {
	Id string
	User user.User
}
type StudentClass struct {
	Id string
	User user.User
}
type HandleClass interface {
	Handle()
}

