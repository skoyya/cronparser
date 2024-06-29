package parser

type MonthParser struct {
	parser
}

const TOTAL_MONTHS  = 12

func NewMonthParser() *MonthParser{
	return &MonthParser{parser {intervalType: "month", floor: 1, ceil: TOTAL_MONTHS, schedule: make([]int, TOKENS_TO_PRINT)}}
}

func (mp *MonthParser) Parse(monthExpression string) error {
	var err error
	mp.schedule, err = mp.parse(monthExpression)
	if err != nil {
		return err
	}
	//Placeholder to add enhance the functionality only specific to MonthParser
	return nil
}

func (mp *MonthParser) PrintIt() {
	mp.printIt()
}
