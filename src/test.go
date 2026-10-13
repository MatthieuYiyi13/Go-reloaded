package reloaded

import "testing"

func TestCap(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello (cap)", "Hello"},
		{"hELLO (cap)", "Hello"},
		{"été (cap)", "Été"},
	}
	for _, test := range tests {
		if got := Cap(test.input); got != test.want {
			t.Errorf("Cap(%q) = %q, attendu %q", test.input, got, test.want)
		}
	}
}
