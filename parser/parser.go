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

func (p WildcardParser) Parse(min, max int) ([]string, error) {
	var values []string
	for i := min; i <= max; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

func (p StepParser) Parse(min, max int) ([]string, error) {
	var values []string
	step := p.Step
	for i := min; i <= max; i += step {
		values = append(values, strconv.Itoa(i))
	}
	return values, nil
}

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

func (p SingleValueParser) Parse(min, max int) ([]string, error) {
	if p.Value < min || p.Value > max {
		return nil, fmt.Errorf("value %d out of range [%d, %d]", p.Value, min, max)
	}
	return []string{strconv.Itoa(p.Value)}, nil
}

// Factory function to get the appropriate parser
func getFieldParser(field string) FieldParser {
	if field == "*" {
		return WildcardParser{}
	}
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, _ := strconv.Atoi(parts[1])
		return StepParser{Step: step}
	}
	if strings.Contains(field, ",") {
		list := strings.Split(field, ",")
		return ListParser{List: list}
	}
	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		return RangeParser{Start: start, End: end}
	}
	num, _ := strconv.Atoi(field)
	return SingleValueParser{Value: num}
}
