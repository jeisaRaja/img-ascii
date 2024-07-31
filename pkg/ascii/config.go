package ascii

import "image"

type ConfigASCII struct {
	letterSet  string
	scaleX     int
	scaleY     int
	brightness float32
	Filepath   string
	img        image.Image
}

const (
	ASCII_SET_1 = "Ñ@#W$9876543210?!abc;:+=-,._ "
	ASCII_SET_2 = "@%#*+=-:. "
	ASCII_SET_3 = "@&%QNW0gB#$DR8mHKAUbOGp4d9hPkqswE2]ayjz/?c*F)J7(Ltv1If{C}r;><=^,':.-` "

	DefaultBrightness float32 = 1.0
	DefaultLetterSet  int     = 1
	DefaultScaleX     int     = 20.0
	DefaultScaleY     int     = 10.0
)

var ASCII_SET = []string{
	ASCII_SET_1,
	ASCII_SET_2,
	ASCII_SET_3,
}

func DefaultConfig() *ConfigASCII {
	return &ConfigASCII{
		letterSet:  ASCII_SET_1,
		scaleX:     DefaultScaleX,
		scaleY:     DefaultScaleY,
		brightness: DefaultBrightness,
	}
}

func (cfg *ConfigASCII) LetterSet(i int) {
	if i > len(ASCII_SET) {
		return
	}
	cfg.letterSet = ASCII_SET[i-1]
}

func (cfg *ConfigASCII) Brightness(i float32) {
	if i > 2 {
		i = 2
	}
	cfg.brightness = i
}

func (cfg *ConfigASCII) ScaleX(i int) {
	cfg.scaleX = i
}

func (cfg *ConfigASCII) ScaleY(i int) {
	cfg.scaleY = i
}
