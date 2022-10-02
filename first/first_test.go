package first_test

import (
	"testing"

	"github.com/tiozafrem/study/first"
)

func TestSubstractMK(t *testing.T) {
	var monthValue []first.MonthValue
	c := first.SubstractMK(100, 50, 2, 2)

	for value := range c {
		monthValue = append(monthValue, value)
	}

	if len(monthValue) != 12 {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return %d value; want 12", len(monthValue))
	}

	firstMonth := first.MonthValue{
		Month: 1,
		Value: 50,
	}
	if monthValue[0] != firstMonth {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return 1 value %v; want %v", monthValue[0], firstMonth)
	}

	lastMonth := first.MonthValue{
		Month: 12,
		Value: 84.30086,
	}
	if monthValue[11] != lastMonth {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return 12 value %v; want %v", monthValue[11], lastMonth)
	}

}

func TestSubstractAdvancedMK(t *testing.T) {
	var monthValue []first.MonthValue
	c := first.SubstractAdvancedMK(100, 50, 2, 2)

	for value := range c {
		monthValue = append(monthValue, value)
	}

	if len(monthValue) != 12 {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return %d value; want 12", len(monthValue))
	}

	firstMonth := first.MonthValue{
		Month: 1,
		Value: 50,
	}
	if monthValue[0] != firstMonth {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return 1 value %v; want %v", monthValue[0], firstMonth)
	}

	lastMonth := first.MonthValue{
		Month: 12,
		Value: 76.30998,
	}
	if monthValue[11] != lastMonth {
		t.Errorf("first.SubstractMK(100, 50, 2, 2) channels return 12 value %v; want %v", monthValue[11], lastMonth)
	}
}

func TestAverageValueYear(t *testing.T) {
	got := first.AverageValueYear(50, 2)
	if got != 56.969837 {
		t.Errorf("first.AverageValueYear(50, 2) = %f; want 56.969837", got)
	}
}
