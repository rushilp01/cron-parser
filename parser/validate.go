package parser

import (
	"cron-parser/constants"
	"fmt"
	"strings"
)

type CronParser struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

// Validate function to parse the fields and return a CronParser struct
/*
Validate function validates the arguments and returns a CronParser struct with all the cron values and expressions
*/
func Validate(fields []string) (CronParser, error) {
	if len(fields) != 6 {
		return CronParser{}, fmt.Errorf("invalid number of fields: expected 6, got %d", len(fields))
	}
	parser, err := getFieldParser(fields[constants.MinuteIndex])
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing minute field: %v", err)
	}
	minuteExpanded, err := parser.Parse(constants.MinuteMinValue, constants.MinuteMaxValue)
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing minute field: %v", err)
	}
	parser, err = getFieldParser(fields[constants.HourIndex])
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing hour field: %v", err)
	}
	hourExpanded, err := parser.Parse(constants.HourMinValue, constants.HourMaxValue)
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing hour field: %v", err)
	}
	parser, err = getFieldParser(fields[constants.DayOfMonthIndex])
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing day of month field: %v", err)
	}
	dayOfMonthExpanded, err := parser.Parse(constants.DayOfMonthMinValue, constants.DayOfMonthMaxValue)
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing day of month field: %v", err)
	}
	parser, err = getFieldParser(fields[constants.MonthIndex])
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing month field: %v", err)
	}
	monthExpanded, err := parser.Parse(constants.MonthMinValue, constants.MonthMaxValue)
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing month field: %v", err)
	}
	parser, err = getFieldParser(fields[constants.DayOfWeekIndex])
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing day of week field: %v", err)
	}
	dayOfWeekExpanded, err := parser.Parse(constants.DayOfWeekMinValue, constants.DayOfWeekMaxValue)
	if err != nil {
		return CronParser{}, fmt.Errorf("error parsing day of week field: %v", err)
	}

	return CronParser{
		Minute:     minuteExpanded,
		Hour:       hourExpanded,
		DayOfMonth: dayOfMonthExpanded,
		Month:      monthExpanded,
		DayOfWeek:  dayOfWeekExpanded,
		Command:    fields[constants.CommandIndex],
	}, nil
}

func (p CronParser) Print() {
	fmt.Printf("%-14s%s\n", "minute", strings.Join(p.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(p.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(p.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(p.Month, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(p.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", p.Command)
}
