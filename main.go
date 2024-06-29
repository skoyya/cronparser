package main
import "fmt"
import "parser"
func main() {
	cp := parser.NewCronParser()
	err := cp.Parse("*/15 0 1,15 * 1-5 my-command")
	if err != nil {
		fmt.Println(err)
		return
	}
	cp.PrintIt()
	return 
}
