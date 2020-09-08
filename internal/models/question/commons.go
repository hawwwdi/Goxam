package question

type JSoner interface {
	JSon() string
}

type Question interface {
	SetId(string)
	GetId() string
	CheckAnswer(...string) bool
	Save() error
	String() string
	JSoner
}

// Dsc is alias for Description
type Dsc struct {
	Text    string
	Picture string
}

func NewDescription(txt , img string) Dsc {
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
