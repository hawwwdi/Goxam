package question

type TestOpt struct {
	Description
	Value bool
}

func NewTestOpt(dsc Description, value bool) TestOpt {
	o := new(TestOpt)
	o.Description = dsc
	o.Value = value
	return *o
}

func (o *TestOpt) Form() string {
	//todo
	return "this is form"
}

func (o *TestOpt) String() string {
	//todo
	return "this is string"
}

type ShortAnswerOpt struct {
	Value []rune
}

func NewShortAnswerOpt(ans []rune) ShortAnswerOpt {
	s := new(ShortAnswerOpt)
	s.Value = ans
	return *s
}

func (o *ShortAnswerOpt) Form() string {
	//todo
	return "this is form"
}

func (o *ShortAnswerOpt) String() string {
	//todo
	return "this is String"
}

