package parser

import (
	"testing"
	"reflect"
)

// TestNewDayOfWeekParser tests the NewDayOfWeekParser function
func TestNewDayOfWeekParser(t *testing.T) {
	dow := NewDayOfWeekParser()
	if dow == nil {
		t.Fatal("Expected DayOfWeekParser instance, got nil")
	}
	if dow.ceil != TOTAL_DAYS_OF_WEEK {
		t.Errorf("Expected ceil to be %d, got %d", TOTAL_DAYS_OF_WEEK, dow.ceil)
	}
	if dow.floor != 1 {
		t.Errorf("Expected floor to be 1, got %d", dow.floor)
	}
	if dow.intervalType != "day of week" {
		t.Errorf("Expected intervalType to be 'day of week', got '%s'", dow.intervalType)
	}
}

// TestValidate tests the Validate method
func TestDayOfWeekParser_Validate(t *testing.T) {
	dow := NewDayOfWeekParser()
	err := dow.Validate("*/2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDayOfWeekParser(t *testing.T) {
	dow := NewDayOfWeekParser()

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/2", []int{1,3,5,7}, false},
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-3", []int{1, 2, 3}, false},
		{"13", nil, true},
		{"*/invalid", nil, true},
		{"1-13", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := dow.Parse(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if err == nil && !reflect.DeepEqual(dow.schedule, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, dow.schedule)
			}
		})
	}
}

