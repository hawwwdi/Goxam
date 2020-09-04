package question

type Question interface {
	CheckAnswer(string) bool
	String() string
	ToForm() string
}



