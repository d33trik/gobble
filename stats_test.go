package gobble_test

import (
	"testing"

	"github.com/d33trik/gobble"
)

func TestAdd(t *testing.T) {
	type input struct {
		stats gobble.Stats
		other gobble.Stats
	}

	tests := map[string]struct {
		input input
		want  gobble.Stats
	}{
		"simple add": {
			input: input{
				stats: gobble.Stats{
					Lines: 5,
					Words: 5,
					Bytes: 24,
				},
				other: gobble.Stats{
					Lines: 1,
					Words: 1,
					Bytes: 4,
				},
			},
			want: gobble.Stats{
				Lines: 6,
				Words: 6,
				Bytes: 28,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.stats
			got.Add(tc.input.other)

			if got != tc.want {
				t.Logf("got: %v, want: %v", got, tc.want)
				t.Fail()
			}
		})
	}
}
