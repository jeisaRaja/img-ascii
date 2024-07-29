package main

import (
	"flag"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/interactive"
	"github.com/jeisaraja/img-ascii/pkg/normal"
)

// var asciiChars = "Ñ@#W$9876543210?!abc;:+=-,._           "
var asciiChars = "@@%#*+=-:.   "

//var asciiChars = "@&%QNW0gB#$DR8mHKAUbOGp4d9hPkqswE2]ayjz/?c*F)J7(Ltv1If{C}r;><=^,':.-`   "

func main() {
	i := flag.Bool("i", false, "interactive mode")
	flag.Parse()
	if *i {
		interactive.Start(os.Stdin, os.Stdout)
	} else {
		normal.NormalMode()
	}
}
