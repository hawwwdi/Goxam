package exam

import (
	"github.com/hawwwdi/Goxam/internal/models/question"
	"github.com/rs/xid"
	"time"
)

type Exam struct {
	question.Description
	id   string
	time time.Duration
	noq  int // noq is alias for Number of question
}

func NewExam(dsc question.Description, time time.Duration) Exam {
	e := new(Exam)
	e.Description = dsc
	e.time = time
	e.id = xid.New().String()
	return *e
}

func (e *Exam) AddQuestion(question question.Question) {
	//todo save question with exam id in database
	e.noq++
}

func (e *Exam) RemQuestion(index int) error {
	//todo remove question from database
	// if question found
	// e.noq--
	return nil
}

func (e *Exam) SetTime(duration time.Duration) {
	e.time = duration
}
