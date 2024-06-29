package parser

type MonthParser struct {
	parser
}

const TOTAL_MONTHS  = 12

func NewMonthParser() *MonthParser{
	return &MonthParser{parser {intervalType: "month", floor: 1, ceil: TOTAL_MONTHS, schedule: make([]int, TOKENS_TO_PRINT)}}
}

//Add Month specific validations here
func (mp *MonthParser) Validate(monthExpression string) error {
	return nil 
}

func (mp *MonthParser) Parse(monthExpression string) error {
	var err error
	mp.schedule, err = mp.parse(monthExpression)
	if err != nil {
		return err
	}
	//Placeholder to add or enhance the functionality only specific to MonthParser
	return nil
}

func (mp *MonthParser) PrintIt() {
	mp.printIt()
}
