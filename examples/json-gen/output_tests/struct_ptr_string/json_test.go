package test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

func Test_Roundtrip(t *testing.T) {
	fz := fuzz.New()
	for i := 0; i < 100; i++ {
		var before T
		var after T

		fz.Fuzz(&before)
		jb, err := Marshal_struct_ptr_string_T(before)
		if err != nil {
			t.Errorf("failed to marshal: %v", err)
		}
		err = Unmarshal_struct_ptr_string_T(jb, &after)
		if err != nil {
			t.Errorf("failed to unmarshal: %v", err)
		}
		if before != after {
			t.Errorf("expected %v, got %v via %q", before, after, string(jb))
		}
	}
}
