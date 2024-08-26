package tests

import (
	"cron-parser/constants"
	"cron-parser/parser"
	"reflect"
	"strconv"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		input    []string
		expected parser.Parser
		err      bool
	}{
		{
			input: []string{"*", "*", "*", "*", "*", "/path/to/command"},
			expected: parser.CronParser{
				Minute:     generateSequence(constants.MinuteMinValue, constants.MinuteMaxValue),
				Hour:       generateSequence(constants.HourMinValue, constants.HourMaxValue),
				DayOfMonth: generateSequence(constants.DayOfMonthMinValue, constants.DayOfMonthMaxValue),
				Month:      generateSequence(constants.MonthMinValue, constants.MonthMaxValue),
				DayOfWeek:  generateSequence(constants.DayOfWeekMinValue, constants.DayOfWeekMaxValue),
				Command:    "/path/to/command",
			},
			err: false,
		},
		{
			input: []string{"*/15", "0", "1,15", "1-5", "*", "/path/to/command"},
			expected: parser.CronParser{
				Minute:     generateStepSequence(constants.MinuteMinValue, constants.MinuteMaxValue, 15),
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5"},
				DayOfWeek:  generateSequence(constants.DayOfWeekMinValue, constants.DayOfWeekMaxValue),
				Command:    "/path/to/command",
			},
			err: false,
		},
		{
			input: []string{"0", "12", "10-20", "2", "1-5", "/path/to/command"},
			expected: parser.CronParser{
				Minute:     []string{"0"},
				Hour:       []string{"12"},
				DayOfMonth: []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"},
				Month:      []string{"2"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    "/path/to/command",
			},
			err: false,
		},
	}

	for _, test := range tests {
		result, err := parser.Validate(test.input)
		if (err != nil) != test.err {
			t.Errorf("Validate(%v) returned error: %v, expected error: %v", test.input, err, test.err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Validate(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestValidateFailure(t *testing.T) {
	tests := []struct {
		input    []string
		expected parser.Parser
		err      bool
	}{
		{
			input:    []string{"*", "*", "*", "*", "*", "/path/to/command", "test/fail"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"60", "*", "*", "*", "*", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"*", "*", "0", "*", "*", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"*", "*", "*", "*", "8", "/path/to/command"},
			expected: nil,
			err:      true,
		},
		{
			input:    []string{"*/2", "*", "*", "*", "0-7", "/path/to/command"},
			expected: nil,
			err:      true,
		},
	}
	for _, test := range tests {
		result, err := parser.Validate(test.input)
		if (err != nil) != test.err {
			t.Errorf("Validate(%v) returned error: %v, expected error: %v", test.input, err, test.err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Validate(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func generateSequence(start, end int) []string {
	var result []string
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func generateStepSequence(start, end, step int) []string {
	var result []string
	for i := start; i <= end; i += step {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
