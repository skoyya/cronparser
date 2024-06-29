package parser

import (
	"testing"
	"reflect"
)

// TestNewMonthParser tests the NewMonthParser function
func TestNewMonthParser(t *testing.T) {
	mp := NewMonthParser()
	if mp == nil {
		t.Fatal("Expected MonthParser instance, got nil")
	}
	if mp.ceil != TOTAL_MONTHS {
		t.Errorf("Expected ceil to be %d, got %d", TOTAL_MONTHS, mp.ceil)
	}
	if mp.floor != 1 {
		t.Errorf("Expected floor to be 1, got %d", mp.floor)
	}
	if mp.intervalType != "month" {
		t.Errorf("Expected intervalType to be 'month', got '%s'", mp.intervalType)
	}
}

// TestValidate tests the Validate method
func TestMonthParser_Validate(t *testing.T) {
	mp := NewMonthParser()
	err := mp.Validate("*/2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestMonthParser(t *testing.T) {
	mp := NewMonthParser()

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/2", []int{1, 3, 5, 7, 9, 11}, false},
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-3", []int{1, 2, 3}, false},
		{"13", nil, true},
		{"*/invalid", nil, true},
		{"1-13", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := mp.Parse(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if err == nil && !reflect.DeepEqual(mp.schedule, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, mp.schedule)
			}
		})
	}
}

