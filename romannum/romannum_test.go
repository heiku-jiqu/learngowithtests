package romannum

import (
	"fmt"
	"testing"
	"testing/quick"
)

var testCases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestPropertyConsecutiveSymbols(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(arabic)
		t.Log("testing", arabic, roman)
		if len(roman) < 4 {
			return true
		}
		for i := 0; i < len(roman)-3; i++ {
			if roman[i] == roman[i+1] && roman[i] == roman[i+2] && roman[i] == roman[i+3] {
				t.Log("i", roman[i], "i+1", roman[i+1], "i+2", roman[i+2])
				return false
			}
		}
		return true
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
