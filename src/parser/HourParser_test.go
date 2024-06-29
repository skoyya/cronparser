package parser

import (
	"testing"
	"reflect"
)

// TestNewHourParser tests the NewHourParser function
func TestNewHourParser(t *testing.T) {
	hp := NewHourParser()
	if hp == nil {
		t.Fatal("Expected HourParser instance, got nil")
	}
	if hp.ceil != TOTAL_HOURS {
		t.Errorf("Expected ceil to be %d, got %d", TOTAL_HOURS, hp.ceil)
	}
	if hp.floor != 0 {
		t.Errorf("Expected floor to be 0, got %d", hp.floor)
	}
	if hp.intervalType != "hour" {
		t.Errorf("Expected intervalType to be 'hour', got '%s'", hp.intervalType)
	}
}

// TestValidate tests the Validate method
func TestHourParser_Validate(t *testing.T) {
	hp := NewHourParser()
	err := hp.Validate("*/2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestHourParser(t *testing.T) {
	hp := NewHourParser()

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/2", []int{0,2,4,6,8,10,12,14,16,18,20,22}, false},
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-3", []int{1, 2, 3}, false},
		{"24", nil, true},
		{"*/invalid", nil, true},
		{"1-24", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := hp.Parse(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if err == nil && !reflect.DeepEqual(hp.schedule, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, hp.schedule)
			}
		})
	}
}

