package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const TOKENS_TO_PRINT = 14

//private to package. Should not be allowed to create an instance of it
//Its like an abstract class in Java.

type parser struct {
	intervalType string
	floor int
	ceil int
	schedule []int
}

//limit the intervals to 14 as mentioned in the execize
func (p *parser) generateIntervals(from int, to int, interval int) []int {
    var intervals []int
    for i := from; i <= to; i += interval {
        intervals = append(intervals, i)
    }
    return intervals
}

func (p *parser) printIt() {
	fmt.Println(p.intervalType, p.schedule)
}

func (p *parser) buildError(expression string) error {
	return fmt.Errorf("invalid %s expression: %s", p.intervalType, expression)
}

func (p *parser) parse(expression string) ([]int, error) {
	if strings.HasPrefix(expression, "*/") {
		interval, err := strconv.Atoi(expression[2:])
		if err != nil || interval >= p.ceil {
			return nil, p.buildError(expression)
		}
		return p.generateIntervals(p.floor, p.ceil, interval), nil
	} else if expression == "*" {
		return p.generateIntervals(p.floor, p.ceil, 1), nil
	} else if strings.Contains(expression, ",") == true {
		tokens := strings.Split(expression,",")
		var intervals []int
		for _, interval := range tokens {
			interval, err := strconv.Atoi(interval)
			if err != nil || interval > p.ceil {
				return nil, p.buildError(expression)
			}
			intervals = append(intervals, interval)
		}
		return intervals, nil
	} else if strings.Contains(expression, "-") {
        tokens := strings.Split(expression,"-")
        if len(tokens) != 2 {
            return nil, p.buildError(expression)
        }
        from, err1 := strconv.Atoi(tokens[0])
        to, err2 := strconv.Atoi(tokens[1])
        if err1 != nil || err2 != nil || from >= p.ceil || to > p.ceil {
            return nil, p.buildError(expression)
        }   
		return p.generateIntervals(from,to, 1), nil
    } else {
		interval, err := strconv.Atoi(expression)
		if err != nil || interval >= p.ceil {
			return nil, p.buildError(expression)
		}
		return []int{interval} , nil
	}

	return nil, p.buildError(expression)
}
