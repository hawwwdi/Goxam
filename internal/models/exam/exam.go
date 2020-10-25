package exam

import (
	"fmt"
	"github.com/hawwwdi/Goxam/internal/models/question"
	"time"
)

type Exam struct {
	question.Dsc
	Id       string
	Time     time.Time
	Duration int // minutes
	Noq      int // noq is alias for Number of question
	Forfeit  int
}

// id in format $ClassId_examIndex
func NewExam(dsc question.Dsc,dateTime time.Time, duration int, id string, forfeit int) Exam {
	e := new(Exam)
	e.Dsc = dsc
	e.Time = dateTime
	e.Duration = duration
	e.Id = id
	e.Forfeit = forfeit
	//todo save the exam
	return *e
}

func (e *Exam) AddQuestion(ques question.Question) {
	var id string
	switch ques.(type) {
	case *question.Test:
		id = fmt.Sprintf("%v_QT%v", e.Id, e.Noq)
	case *question.ShortAnswer:
		id = fmt.Sprintf("%v_QS%v", e.Id, e.Noq)
	}
	ques.SetId(id)
	_ = ques.Save()
	//todo handle above error
	e.Noq++
}

func (e *Exam) RemQuestion(id string) error {
	//todo remove question from database
	// if question found
	// rem question
	// e.noq--
	// endif
	return nil
}

/*func (e *Exam) SetTime(duration time.Duration) {
	e.time = duration
}

func (e *Exam) Id () string {
	return e.id
}*/

func (e *Exam) String() string {
	return "this is string"
}
