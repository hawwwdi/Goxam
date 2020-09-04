package question

import "image"

type Option struct {
	Description
	Value bool
}

func NewOption(txt []rune, img image.Image, value bool) Option {
	o := new(Option)
	o.Text = txt
	o.Picture = img
	o.Value = value
	return *o
}

func (o *Option) Form() string {
	//todo
	return "this is form"
}

func (o *Option) String() string {
	//todo
	return "this is string"
}
