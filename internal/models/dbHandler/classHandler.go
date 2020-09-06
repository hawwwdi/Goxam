package dbHandler

import (
	"database/sql"
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/user"
)

func SaveClass(db *sql.DB , user user.User , id string)  {
	stmt, err := db.Prepare(`INSERT INTO classes VALUES (?,?);`)
	errHandler(err)
	defer stmt.Close()
	r, err := stmt.Exec(id , user.Email)
	errHandler(err)
	ro, err := r.RowsAffected()
	errHandler(err)
	fmt.Println("INSERTED RECORD TO Classes", ro)
}
