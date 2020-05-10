package color

import "fmt"

var c = New(Plain)

// Style switches style
func Style(style int) {
	c.Style(style)
}

// Black returns black plain text
func Black(s string) string {
	return c.Black(s)
}

// Red returns red plain text
func Red(s string) string {
	return c.Red(s)
}

// Green returns green plain text
func Green(s string) string {
	return c.Green(s)
}

// Orange returns orange plain text
func Orange(s string) string {
	return c.Orange(s)
}

// Blue returns blue plain text
func Blue(s string) string {
	return c.Blue(s)
}

// Purple returns purple plain text
func Purple(s string) string {
	return c.Purple(s)
}

// Cyan returns cyan plain text
func Cyan(s string) string {
	return c.Cyan(s)
}

// Gray returns gray plain text
func Gray(s string) string {
	return c.Gray(s)
}

// White returns white plain text
func White(s string) string {
	var style = Plain
	if c.style != Plain {
		style = c.style
	}
	return fmt.Sprintf(Template, style, gray, s)
}
