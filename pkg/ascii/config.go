package ascii

type ConfigASCII struct {
	LetterSet  string
	ScaleX     int
	ScaleY     int
	Contrast   float32
	Brightness float32
}

const (
	ASCII_SET_1 = "Ñ@#W$9876543210?!abc;:+=-,._ "
	ASCII_SET_2 = "@@%#*+=-:. "
	ASCII_SET_3 = "@&%QNW0gB#$DR8mHKAUbOGp4d9hPkqswE2]ayjz/?c*F)J7(Ltv1If{C}r;><=^,':.-` "
)

func DefaultConfig() *ConfigASCII {
	return &ConfigASCII{
		LetterSet:  ASCII_SET_3,
		ScaleX:     20,
		ScaleY:     10,
    Brightness: 2.5,
	}
}
