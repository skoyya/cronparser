package parser
import "fmt"

type CommandParser struct {
	command string
}

func NewCommandParser() *CommandParser{
	return &CommandParser{""}
}

func (mp *CommandParser) Parse(commandExpression string) error {
	  mp.command = commandExpression
	  return nil
}

func (mp *CommandParser) PrintIt() {
	fmt.Println("command", mp.command)
}
