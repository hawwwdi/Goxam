package exam

import (
	"github.com/hawwwdi/Goxam/internal/models/question"
	"time"
)

type Exam struct {
	question.Dsc
	id   string
	time time.Duration
	noq  int // noq is alias for Number of question
	Forfeit int
}

// id in format $ClassId_examIndex
func NewExam(dsc question.Dsc, time time.Duration, id string, forfeit int) Exam {
	e := new(Exam)
	e.Dsc = dsc
	e.time = time
	e.id = id
	e.Forfeit = forfeit
	return *e
}

func (e *Exam) AddQuestion(question question.Question) {
	//todo save question with exam id in database
	e.noq++
}

func (e *Exam) RemQuestion(index int) error {
	//todo remove question from database
	// if question found
	// rem question
	// e.noq--
	// endif
	return nil
}

func (e *Exam) SetTime(duration time.Duration) {
	e.time = duration
}

func (e *Exam) Id () string {
	return e.id
}

func (e *Exam) String() string {
	return "this is string"
}
