package musgo_uint

import "testing"

func TestMusgoUint(t *testing.T) {
	var n uint = 5
	buf := make([]byte, Uint(n).SizeMUS())
	Uint(n).MarshalMUS(buf)

	var an uint
	(*Uint)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
