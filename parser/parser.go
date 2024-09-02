package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// FieldParser interface for parsing fields
type FieldParser interface {
	Parse(min, max int) ([]string, error)
}

// WildcardParser for '*'
type WildcardParser struct{}

// StepParser for fields with '/'
type StepParser struct {
	Step int
}

// ListParser for fields with ','
type ListParser struct {
	List []string
}

// RangeParser for fields with '-'
type RangeParser struct {
	Start, End int
}

// SingleValueParser for single values
type SingleValueParser struct {
	Value int
}

// Parse
/**
 * Parses a * pattern in the cron expression.
 *
 * @param min   The minimum allowed value for the expression.
 * @param max   The maximum allowed value for the expression.
 * @return A list of integers converted to string representing the parsed values for the * pattern.
 */
func (p WildcardParser) Parse(min, max int) ([]string, error) {
	var values []string
	for i := min; i <= max; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

// Parse
/**
 * Parses a step(/) pattern in the cron expression.
 *
 * @param min   The minimum allowed value for the expression.
 * @param max   The maximum allowed value for the expression.
 * @return A list of integers converted to string representing the parsed values for the / pattern.
 */

func (p StepParser) Parse(min, max int) ([]string, error) {
	var values []string
	step := p.Step
	for i := min; i <= max; i += step {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

// Parse
/**
 * Parses a list pattern in the cron expression.
 *
 * @param min   The minimum allowed value for the expression.
 * @param max   The maximum allowed value for the expression.
 * @return A list of integers converted to string representing the parsed values for the , pattern.
 */
func (p ListParser) Parse(min, max int) ([]string, error) {
	var values []string
	for _, v := range p.List {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		if num < min || num > max {
			return nil, fmt.Errorf("value %d out of range [%d, %d]", num, min, max)
		}
		values = append(values, strconv.Itoa(num))
	}
	return values, nil
}

// Parse
/**
 * Parses a range pattern in the cron expression.
 *
 * @param min   The minimum allowed value for the expression.
 * @param max   The maximum allowed value for the expression.
 * @return A list of integers converted to string representing the parsed values for the - pattern.
 */
func (p RangeParser) Parse(min, max int) ([]string, error) {
	var values []string
	if p.Start < min || p.End > max {
		return nil, fmt.Errorf("range %d-%d out of range [%d, %d]", p.Start, p.End, min, max)
	}
	for i := p.Start; i <= p.End; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

// Parse
/**
 * Parses a single value pattern in the cron expression.
 *
 * @param min   The minimum allowed value for the expression.
 * @param max   The maximum allowed value for the expression.
 * @return A list of integer converted to string representing the parsed values for the single value pattern.
 */
func (p SingleValueParser) Parse(min, max int) ([]string, error) {
	if p.Value < min || p.Value > max {
		return nil, fmt.Errorf("value %d out of range [%d, %d]", p.Value, min, max)
	}
	return []string{strconv.Itoa(p.Value)}, nil
}

// Factory function to get the appropriate parser
func getFieldParser(field string) (FieldParser, error) {
	if field == "*" {
		return WildcardParser{}, nil
	}
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		return StepParser{Step: step}, nil
	}
	if strings.Contains(field, ",") {
		list := strings.Split(field, ",")
		return ListParser{List: list}, nil
	}
	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return nil, err
		}
		return RangeParser{Start: start, End: end}, nil
	}
	num, err := strconv.Atoi(field)
	if err != nil {
		return nil, err
	}
	return SingleValueParser{Value: num}, nil
}
