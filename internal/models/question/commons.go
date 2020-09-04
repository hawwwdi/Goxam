package question

import "image"

type Question interface {
	CheckAnswer(...string) bool
	String() string
	Form() string
}

type Description struct {
	Text []rune
	Picture image.Image
}


