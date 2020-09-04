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
	Text []rune
	Picture image.Image
}


