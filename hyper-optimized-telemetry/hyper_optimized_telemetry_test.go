package hotelemetry_test

import (
	hotlm "hotelemetry"
	"testing"
)

func TestToBuffer(t *testing.T) {
	t.Parallel()
	for name, tc := range toBufferTT {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := hotlm.ToBuffer(tc.value); got != tc.want {
				t.Fatalf("\n\twant: %+v\n\tgot:  %+v", tc.want, got)
			}
		})
	}
}

func TestFromBuffer(t *testing.T) {
	t.Parallel()
	for name, tc := range fromBufferTT {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := hotlm.FromBuffer(tc.buf); got != tc.want {
				t.Fatalf("\n\twant: %+v\n\tgot:  %+v", tc.want, got)
			}
		})
	}
}
