package question

import "image"

type Former interface {
	Form() string
}

type Question interface {
	CheckAnswer(...string) bool
	String() string
	Former
}

type Description struct {
	Text    string
	Picture image.Image
}

func NewDescription(txt string, img image.Image) Description {
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
