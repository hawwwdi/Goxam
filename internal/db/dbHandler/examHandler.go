package dbHandler

import "github.com/hawwwdi/Goxam/internal/models/exam"

func SaveExam(exm *exam.Exam, classId string) error {
	query := `INSERT INTO exams VALUES (?,?,?,?,?,?,?,?);`
	stmt, err := db.Prepare(query)
	//errHandler(err)
	if err != nil {
		return err
	}
	_, err1 := stmt.Exec(
		classId,
		exm.Id,
		exm.DateTime,
		exm.Time,
		exm.Forfeit,
		exm.Noq,
		exm.Dsc.Text,
		exm.Id,
	)
	//todo change the image id
	//errHandler(err1)
	if err1 != nil {
		return err1
	}
	return nil
}
