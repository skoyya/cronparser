package parser

type ICronParser interface {
	Parse(expression string) error
	PrintIt() 
}

