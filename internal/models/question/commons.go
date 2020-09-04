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
	Text    []rune
	Picture image.Image
}

func NewDescription(txt []rune, img image.Image) Description {
	d := new(Description)
	d.Text = txt
	d.Picture = img
	return *d
}
