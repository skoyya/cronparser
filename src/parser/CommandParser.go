package parser
import "fmt"

type CommandParser struct {
	command string
}

func NewCommandParser() *CommandParser{
	return &CommandParser{""}
}

func (cp *CommandParser) Parse(commandExpression string) error {
	  cp.command = commandExpression
	  return nil
}

//Add Command specific validations here
func (cp *CommandParser) Validate(commandExpression string) error {
    return nil
}

func (cp *CommandParser) PrintIt() {
	fmt.Println("command", cp.command)
}
