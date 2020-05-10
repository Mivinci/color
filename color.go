package color

import "fmt"

// Color codes
const (
	black = 30 + iota
	red
	green
	orange
	blue
	purple
	cyan
	gray
)

// Style codes
const (
	Plain = iota
	Bright
	Dark
	Italic
	Underlined
)

// Template
const (
	Template = "\033[%d;%dm%s\033[0m"
)

// Color is color handler
type Color struct {
	style int
}

// New new a color handler
func New(style int) *Color { return &Color{style: style} }

// Style sets style
func (c *Color) Style(style int) {
	c.style = style
}

// Black returns black text
func (c *Color) Black(s string) string {
	return fmt.Sprintf(Template, c.style, black, s)
}

// Red returns red text
func (c *Color) Red(s string) string {
	return fmt.Sprintf(Template, c.style, red, s)
}

// Green returns green text
func (c *Color) Green(s string) string {
	return fmt.Sprintf(Template, c.style, green, s)
}

// Orange returns orange text
func (c *Color) Orange(s string) string {
	return fmt.Sprintf(Template, c.style, orange, s)
}

// Blue returns blue text
func (c *Color) Blue(s string) string {
	return fmt.Sprintf(Template, c.style, blue, s)
}

// Purple returns purple text
func (c *Color) Purple(s string) string {
	return fmt.Sprintf(Template, c.style, purple, s)
}

// Cyan returns cyan text
func (c *Color) Cyan(s string) string {
	return fmt.Sprintf(Template, c.style, cyan, s)
}

// Gray returns gray text
func (c *Color) Gray(s string) string {
	return fmt.Sprintf(Template, c.style, gray, s)
}
