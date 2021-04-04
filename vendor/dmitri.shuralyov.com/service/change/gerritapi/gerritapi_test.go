package gerritapi

import (
	"testing"
)

func TestParseMessage(t *testing.T) {
	tests := []struct {
		in         string
		wantLabels string
		wantBody   string
	}{
		{
			in:         "Patch Set 2: Code-Review+2",
			wantLabels: "Code-Review+2",
			wantBody:   "",
		},
		{
			in:         "Patch Set 3: Run-TryBot+1 Code-Review+2",
			wantLabels: "Run-TryBot+1 Code-Review+2",
			wantBody:   "",
		},
		{
			in:         "Patch Set 2: Code-Review+2\n\nThanks.",
			wantLabels: "Code-Review+2",
			wantBody:   "Thanks.",
		},
		{
			in:         "Patch Set 1:\n\nFirst contribution — trying to get my feet wet. Please review.",
			wantLabels: "",
			wantBody:   "First contribution — trying to get my feet wet. Please review.",
		},
	}
	for i, tc := range tests {
		gotLabels, gotBody, ok := parseMessage(tc.in)
		if !ok {
			t.Fatalf("%d: not ok", i)
		}
		if gotLabels != tc.wantLabels {
			t.Errorf("%d: got labels: %q, want: %q", i, gotLabels, tc.wantLabels)
		}
		if gotBody != tc.wantBody {
			t.Errorf("%d: got body: %q, want: %q", i, gotBody, tc.wantBody)
		}
	}
}
