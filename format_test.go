package adif

import "testing"

func Test_normalizeFreqMHz(t *testing.T) {
	cases := map[string]string{
		"7.050.000": "7.050",
		"14.074":    "14.074",
		"144.390":   "144.390",
		"144.":      "144",
		"":          "",
	}
	for in, want := range cases {
		got := normalizeFreqMHz(in)
		if got != want {
			t.Fatalf("normalizeFreqMHz(%q) = %q; want %q", in, got, want)
		}
	}
}
