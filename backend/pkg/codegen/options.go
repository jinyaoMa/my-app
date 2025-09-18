package codegen

const (
	Digits     string = "0123456789"
	HexDigits  string = Digits + "abcdefABCDEF"
	LHexDigits string = Digits + "abcdef"
	UHexDigits string = Digits + "ABCDEF"
	Letters    string = "abcdefghijklmnopqrstuvwxyz"
	ULetters   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func DefaultOptions() Options {
	return Options{
		Characters: Digits + Letters + ULetters,
	}
}

type Options struct {
	Characters string `json:"characters"`
}
