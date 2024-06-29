package parser

import (
	"testing"
	"reflect"
)

// TestNewDayOfMonthParser tests the NewDayOfMonthParser function
func TestNewDayOfMonthParser(t *testing.T) {
	dom := NewDayOfMonthParser()
	if dom == nil {
		t.Fatal("Expected DayOfMonthParser instance, got nil")
	}
	if dom.ceil != TOTAL_DAYS {
		t.Errorf("Expected ceil to be %d, got %d", TOTAL_DAYS, dom.ceil)
	}
	if dom.floor != 1 {
		t.Errorf("Expected floor to be 1, got %d", dom.floor)
	}
	if dom.intervalType != "day of month" {
		t.Errorf("Expected intervalType to be 'day of month', got '%s'", dom.intervalType)
	}
}

// TestValidate tests the Validate method
func TestDayOfMonthParser_Validate(t *testing.T) {
	dom := NewDayOfMonthParser()
	err := dom.Validate("*/2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDayOfMonthParser(t *testing.T) {
	dom := NewDayOfMonthParser()

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/2", []int{1,3,5,7,9,11,13,15,17,19,21,23,25,27,29,31}, false},
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-3", []int{1, 2, 3}, false},
		{"32", nil, true},
		{"*/invalid", nil, true},
		{"1-35", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := dom.Parse(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if err == nil && !reflect.DeepEqual(dom.schedule, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, dom.schedule)
			}
		})
	}
}

