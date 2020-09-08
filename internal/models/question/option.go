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

func (o *TestOpt) JSon() string {
	//todo
	return "this is json"
}

func (o *TestOpt) String() string {
	//todo
	return "this is string"
}

type ShortAnswerOpt struct {
	Value string
}

func NewShortAnswerOpt(ans string) ShortAnswerOpt {
	s := new(ShortAnswerOpt)
	s.Value = ans
	return *s
}

func (o *ShortAnswerOpt) JSon() string {
	//todo
	return "this is json"
}

func (o *ShortAnswerOpt) String() string {
	//todo
	return "this is String"
}

