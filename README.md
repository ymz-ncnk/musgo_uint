# MusGo_uint
Provides serialization of the Golang's `uint` type.

# How to use
Simply cast `uint` to `musgo_uint.Uint`:
```go
	var n uint = 5
	// Marshal
	buf := make([]byte, musgo_uint.Uint(n).SizeMUS())
	musgo_uint.Uint(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_uint.Uint)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

