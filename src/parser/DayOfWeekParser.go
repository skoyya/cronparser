package parser

type DayOfWeekParser struct {
	parser
}

const TOTAL_DAYS_OF_WEEK  = 7

func NewDayOfWeekParser() *DayOfWeekParser{
	return &DayOfWeekParser{parser {intervalType: "day of week", floor: 1, ceil: TOTAL_DAYS_OF_WEEK, schedule: make([]int, TOKENS_TO_PRINT)}}
}

func (dow *DayOfWeekParser) Parse(dayOfWeekExpression string) error {
	var err error
	dow.schedule, err = dow.parse(dayOfWeekExpression)
	if err != nil {
		return err
	}
	//Placeholder to add enhance the functionality only specific to DayOfWeekParser
	return nil
}

func (dow *DayOfWeekParser) PrintIt() {
	dow.printIt()
}
