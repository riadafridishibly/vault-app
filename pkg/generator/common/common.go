package common

const (
	LowerAscii = "abcdefghijklmnopqrstuvwxyz"
	UpperAscii = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers    = "012345678"
	Symbols    = "`~!@#$%^&*()_+-\\'\":;?.<,>/"
	Spaces     = " "
	AllChars   = LowerAscii + UpperAscii + Numbers + Symbols + Spaces
)

type Options struct {
	AllowRepeat    bool
	IncludeSymbols bool
	IncludeSpace   bool
	Length         int
}
