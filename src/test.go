package reloaded

import "testing"

func TestUp(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"hello (up)", "HELLO"},
		{"hello world (up)", "hello WORLD"},
		{"hello (up) world (up)", "HELLO WORLD"},
		{"été (up)", "ÉTÉ"},
		{"(up) hello", "hello"},
		{"hello (up) (up)", "HELLO"},
		{"hello world", "hello world"},
		{"", ""},
	}

	for _, test := range tests {
		if got := Up(test.input); got != test.want {
			t.Errorf("Up(%q) = %q, attendu %q", test.input, got, test.want)
		}
	}
}
