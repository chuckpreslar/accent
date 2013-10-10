// Package accent provides color and decoration accents
// for terminal output.
package accent

import "fmt"

type color [2]uint8

var (
	White   = color{37, 39}
	Grey    = color{90, 39}
	Black   = color{30, 39}
	Blue    = color{34, 39}
	Cyan    = color{36, 39}
	Green   = color{32, 39}
	Magenta = color{35, 39}
	Red     = color{31, 39}
	Yellow  = color{33, 39}
)

type decoration [2]uint8

var (
	Bold          = decoration{1, 22}
	Italic        = decoration{3, 23}
	Underline     = decoration{4, 24}
	Strikethrough = decoration{9, 29}
)

type accentuator interface {
	// accentuate accepts an interface and returns
	// an accented string.
	accentuate(interface{}) string
}

// Implements Stringer interface for color type.
func (c color) String() string {
	return fmt.Sprintf("\x1B[%dm%%v\x1B[%dm", c[0], c[1])
}

// accentuate implements the accentuator interface
// for the color type.
func (c color) accentuate(input interface{}) string {
	return fmt.Sprintf(c.String(), input)
}

// Implements Stringer interface for decoration type.
func (d decoration) String() string {
	return fmt.Sprintf("\x1B[%dm%%v\x1B[%dm", d[0], d[1])
}

// accentuate implements the accentuator interface
// for the decoration type.
func (d decoration) accentuate(input interface{}) string {
	return fmt.Sprintf(d.String(), input)
}

// Message formats a string based on the input parameter
// accented with the provided accentuators and returning
// the resulting string.
func Message(input interface{}, accents ...accentuator) string {
	for i := 0; i < len(accents); i++ {
		input = accents[i].accentuate(input)
	}

	return input.(string)
}
