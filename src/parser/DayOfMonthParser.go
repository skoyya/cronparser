package parser

type DayOfMonthParser struct {
	parser
}

const TOTAL_DAYS  = 31

func NewDayOfMonthParser() *DayOfMonthParser{
	return &DayOfMonthParser{parser {intervalType: "day of month", floor: 1, ceil: TOTAL_DAYS, schedule: make([]int, TOKENS_TO_PRINT)}}
}

func (dom *DayOfMonthParser) Parse(dayOfMonthExpression string) error {
	var err error
	dom.schedule, err = dom.parse(dayOfMonthExpression)
	if err != nil {
		return err
	}
	//Placeholder to add enhance the functionality only specific to DayOfMonthParser
	return nil
}

func (dom *DayOfMonthParser) PrintIt() {
	dom.printIt()
}
