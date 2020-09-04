package question

type test struct {
	MultipleChoice bool
}

func (t *test) CheckAnswer(answer ...string) bool {
	//todo
	return true
}

func (t *test) Form () string {
	//todo
	return "this is form"
}

func (t *test) String() string {
	//todo
	return "this is string"
}
