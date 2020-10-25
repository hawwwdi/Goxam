package dbHandler

import (
	"github.com/hawwwdi/Goxam/internal/models/exam"
	"github.com/hawwwdi/Goxam/internal/models/question"
	"testing"
	"time"
)

func TestSaveExam(t *testing.T) {
	dsc := question.NewDescription(`this is a test`, "test")
	exm := exam.NewExam(dsc, time.Now(), time.Hour, "123", 3)
	if err := SaveExam(&exm, "test"); err != nil {
		t.Error(err)
	}
	err := CloseDB()
	errHandler(err)
}