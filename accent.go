package accent

import "fmt"

type color uint8

const (
	White color = iota
	Grey
	Black
	Blue
	Cyan
	Green
	Magenta
	Red
	Yellow
)

type decoration uint8

const (
	Bold decoration = iota
	Italic
	Underline
	Strikethrough
)

type accentuator interface {
	accentuate(interface{}) string
}

var colors = map[color]string{
	White:   "\x1B[37m%v\x1B[39m",
	Grey:    "\x1B[90m%v\x1B[39m",
	Black:   "\x1B[30m%v\x1B[39m",
	Blue:    "\x1B[34m%v\x1B[39m",
	Cyan:    "\x1B[36m%v\x1B[39m",
	Green:   "\x1B[32m%v\x1B[39m",
	Magenta: "\x1B[35m%v\x1B[39m",
	Red:     "\x1B[31m%v\x1B[39m",
	Yellow:  "\x1B[33m%v\x1B[39m",
}

var decorations = map[decoration]string{
	Bold:          "\x1B[1m%v\x1B[22m",
	Italic:        "\x1B[3m%v\x1B[23m'",
	Underline:     "\x1B[4m%v\x1B[24m",
	Strikethrough: "\x1B[9m%v\x1B[29m",
}

func (c color) accentuate(input interface{}) string {
	return fmt.Sprintf(colors[c], input)
}

func (d decoration) accentuate(input interface{}) string {
	return fmt.Sprintf(decorations[d], input)
}

func Message(input interface{}, accents ...accentuator) string {

	for i := 0; i < len(accents); i++ {
		input = accents[i].accentuate(input)
	}

	return input.(string)
}
