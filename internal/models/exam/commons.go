package exam

import (
	"github.com/gofrs/uuid"
	"github.com/hawwwdi/Goxam/internal/models/question"
	"time"
)

type Exam struct {
	question.Description
	id uuid.UUID
	time time.Duration
}

func NewExam(dsc question.Description, time time.Duration) Exam {
	e := new(Exam)
	e.Description = dsc
	e.time = time
	e.id = uuid.Must(uuid.NewV1())
	return *e
}

func (e *Exam) AddQuestion(question question.Question) {

}

