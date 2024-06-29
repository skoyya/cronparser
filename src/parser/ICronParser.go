package parser

type ICronParser interface {
	Parse(expression string) error
	PrintIt() 
	Validate(expression string) error
}

