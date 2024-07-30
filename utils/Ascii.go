package utils

type Ascii struct {
	Input     string
	SubString string
	Color     string
	Banner    string
}

func NewAscii() *Ascii {
	return &Ascii{
		Banner: "standard.txt",
	}
}
