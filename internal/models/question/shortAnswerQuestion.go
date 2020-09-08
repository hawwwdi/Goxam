package question

import "strings"

type ShortAnswer struct {
	Dsc
	Blanks []ShortAnswerOpt
}

func NewShortAnswer(dsc Dsc) ShortAnswer {
	s := new(ShortAnswer)
	s.Dsc = dsc
	s.Blanks = make([]ShortAnswerOpt, 0, 1)
	return *s
}

func (a *ShortAnswer) AddOption(opt ...ShortAnswerOpt) {
	a.Blanks = append(a.Blanks, opt...)
}

func (a *ShortAnswer) CheckAnswer(answer ...string) bool {
	for i, curr := range answer {
		if !strings.EqualFold(a.Blanks[i].Value, curr) {
			return false
		}
	}
	return true
}


func (a *ShortAnswer) JSon() string {
	//todo
	return "this is form"
}

func (a *ShortAnswer) String() string {
	//todo
	return "this is string"
}
