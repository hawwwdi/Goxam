package question

import "github.com/hawwwdi/Goxam/internal/imageMangement"

type Kind int
// QTest for test type
// QSAns for short answer question
const (
	QTest = iota
	QSAns
)

type Question interface {
	SetId(string)
	GetId() string
	CheckAnswer(...string) bool
	Save() error
	String() string
}

// Dsc is alias for Description
type Dsc struct {
	Text    string
	Picture imageMangement.Image
}

func NewDescription(txt string , img imageMangement.Image) Dsc {
	d := new(Dsc)
	d.Text = txt
	d.Picture = img
	return *d
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
