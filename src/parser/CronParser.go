package parser
import (
	"fmt"
	"strings"
)
type CronParser struct {
	cronParsers []ICronParser
}

func NewCronParser() *CronParser{
	parsers := make([]ICronParser, 6)	
	parsers[0] = NewMinuteParser()
	parsers[1] = NewHourParser()
	parsers[2] = NewDayOfMonthParser()
	parsers[3] = NewMonthParser()
	parsers[4] = NewDayOfWeekParser()
	parsers[5] = NewCommandParser()
	return &CronParser{parsers}
}

func (cp *CronParser) Parse(expression string) error {
	tokens := strings.Fields(expression)
	if len(tokens) != 6 {
		return fmt.Errorf("invalid cron expression")
	}

	for idx, token := range tokens {
		err := cp.cronParsers[idx].Parse(token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cp *CronParser) PrintIt() {
	for _, parser := range cp.cronParsers {
		parser.PrintIt()
	}
}
