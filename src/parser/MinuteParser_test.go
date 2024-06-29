package parser

import (
	"testing"
	"reflect"
)

// TestNewMinuteParser tests the NewMinuteParser function
func TestNewMinuteParser(t *testing.T) {
	mp := NewMinuteParser()
	if mp == nil {
		t.Fatal("Expected MinuteParser instance, got nil")
	}
	if mp.ceil != TOTAL_MINUTES {
		t.Errorf("Expected ceil to be %d, got %d", TOTAL_MINUTES, mp.ceil)
	}
	if mp.floor != 0 {
		t.Errorf("Expected floor to be 1, got %d", mp.floor)
	}
	if mp.intervalType != "minute" {
		t.Errorf("Expected intervalType to be 'minute', got '%s'", mp.intervalType)
	}
}

// TestValidate tests the Validate method
func Test_MinuteParser_Validate(t *testing.T) {
	mp := NewMinuteParser()
	err := mp.Validate("*/2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestMinuteParser(t *testing.T) {
	mp := NewMinuteParser()

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/15", []int{0, 15, 30, 45}, false},
        {"*/10", []int{0, 10, 20, 30, 40, 50}, false},
        {"*", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59}, false},
        {"1,2,3", []int{1, 2, 3}, false},
        {"1-5", []int{1, 2, 3, 4, 5}, false},
		{"70", nil, true},
		{"*/invalid", nil, true},
		{"1-60", nil, true},
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

