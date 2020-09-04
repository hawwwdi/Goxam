package question

type Option struct {
	Description
	Value bool
}

func NewOption(dsc Description, value bool) Option {
	o := new(Option)
	o.Description = dsc
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
