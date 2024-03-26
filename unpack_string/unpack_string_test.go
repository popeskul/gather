package unpack_string

import "testing"

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"3abc", "", true},
		{"45", "", true},
		{"aaa10b", "", true},
		{"aaa0b", "aab", false},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`qw\ne`, "", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := UnpackString(test.input)
			if (err != nil) != test.err {
				t.Errorf("expected error %v, got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("expected %q, got %q", test.expected, result)
			}
		})
	}
}
