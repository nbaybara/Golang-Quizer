package cmd

import "github.com/fatih/color"

//Here is for Color methods

func resetColor() {

	color.Set(color.Reset)
}

func italicFont() {

	color.Set(color.Italic)
}

func textUnderline() {

	color.Set(color.Underline)
}

func boldfont() {

	color.Set(color.Bold)
}

func greenColor() {

	color.Set(color.FgHiGreen)
}

func blueColor() {

	color.Set(color.FgBlue)
}

func redColor() {

	color.Set(color.FgRed)
}

func magentaColor() {

	color.Set(color.FgMagenta)
}

func cyanColor() {

	color.Set(color.FgHiCyan)
}

func yellowColor() {

	color.Set(color.FgYellow)
}

func whiteColor() {

	color.Set(color.FgWhite)
}
