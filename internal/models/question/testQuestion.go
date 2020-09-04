package question

type Test struct {
	Dsc
	MultipleChoice bool
	Options []Option
}

func NewTest(dsc Description, mc bool) Test {
	t := new(Test)
	t.Description = dsc
	t.MultipleChoice = mc
	t.Options = make([]Option, 0, 2)
	return *t
}

func (t *Test) CheckAnswer(answer ...string) bool {
	//todo
	return true
}

func (t *Test) Form () string {
	//todo
	return "this is form"
}

func (t *Test) String() string {
	//todo
	return "this is string"
}
