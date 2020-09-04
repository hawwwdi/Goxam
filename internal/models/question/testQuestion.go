package question

type Test struct {
	MultipleChoice bool
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
