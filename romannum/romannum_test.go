package romannum

import "testing"

func TestRomanNumerals(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to I", 2, "II"},
		{"3 converts to I", 3, "III"},
		{"4 converts to I", 4, "IV"},
		{"5 converts to I", 5, "V"},
		{"6 converts to I", 6, "VI"},
		{"9 converts to I", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertToRoman(test.input)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
