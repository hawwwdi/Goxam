package question

import "strconv"

type Test struct {
	Dsc
	Id string
	MultipleChoice bool
	Options []TestOpt
}

func NewTest(dsc Dsc, id string, mc bool) Test {
	t := new(Test)
	t.Dsc = dsc
	t.Id = id
	t.MultipleChoice = mc
	t.Options = make([]TestOpt, 0, 2)
	return *t
}

func (t *Test) AddOption(opt ...TestOpt) {
	t.Options = append(t.Options, opt...)
}

func (t *Test) CheckAnswer(answer ...string) bool {
	for _, ans := range answer {
		index, err := strconv.Atoi(ans)
		handleError(err)
		if !t.Options[index].Value {
			return false
		}
	}
	return true
}

func (t *Test) JSon() string {
	//todo
	return "this is form"
}

func (t *Test) String() string {
	//todo
	return "this is string"
}


func (t *Test) Save() error {
	//todo
	return nil
}