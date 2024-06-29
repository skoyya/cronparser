package parser

type MinuteParser struct {
	parser
}

const TOTAL_MINUTES  = 59

func NewMinuteParser() *MinuteParser{
	return &MinuteParser{parser {intervalType: "minute", floor: 0, ceil: TOTAL_MINUTES, schedule: make([]int, TOKENS_TO_PRINT)}}
}

func (mp *MinuteParser) Parse(minuteExpression string) error {
	var err error
	mp.schedule, err = mp.parse(minuteExpression)
	if err != nil {
		return err
	}
	//Placeholder to add enhance the functionality only specific to MinuteParser
	return nil
}

func (mp *MinuteParser) PrintIt() {
	mp.printIt()
}
