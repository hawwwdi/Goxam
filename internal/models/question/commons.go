package question

type JSoner interface {
	JSon() string
}

type Question interface {
	CheckAnswer(...string) bool
	String() string
	JSoner
}

type Description struct {
	Text    string
	Picture string
}

func NewDescription(txt , img string) Description {
	d := new(Description)
	d.Text = txt
	d.Picture = img
	return *d
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
