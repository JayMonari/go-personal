package hotlm_test

import (
	hotlm "hotelemetry"
	"testing"
)

func TestToBuffer(t *testing.T) {
	for _, tc := range toBufferTT {
		t.Run(tc.name, func(t *testing.T) {
			if got := hotlm.ToBuffer(tc.value); got != tc.want {
				t.Errorf("\n\twant: %+v\n\tgot:  %+v", tc.want, got)
			}
		})
	}
}

func TestFromBuffer(t *testing.T) {
	for _, tc := range fromBufferTT {
		t.Run(tc.name, func(t *testing.T) {
			if got := hotlm.FromBuffer(tc.buf); got != tc.want {
				t.Errorf("\n\twant: %+v\n\tgot:  %+v", tc.want, got)
			}
		})
	}
}
