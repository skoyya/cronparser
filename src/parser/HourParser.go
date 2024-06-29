package parser

type HourParser struct {
	parser
}

const TOTAL_HOURS = 23

func NewHourParser() *HourParser{
	return &HourParser{parser {intervalType: "hour", floor: 0, ceil: TOTAL_HOURS, schedule: make([]int, TOKENS_TO_PRINT)}}
}

func (hp *HourParser) Parse(hourExpression string) error {
	var err error
	hp.schedule, err = hp.parse(hourExpression)
	if err != nil {
		return err
	}
	//Placeholder to add enhance the functionality only specific to HourParser
	return nil
}

func (hp *HourParser) PrintIt() {
	hp.printIt()
}
