package main

import (
	"fmt"

	"github.com/mivinci/color"
)

func main() {
	fmt.Println(color.White("white"))
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Orange("orange"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.Gray("gray"))

	color.Style(color.Bright)
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Orange("orange"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.Gray("gray"))

	color.Style(color.Italic)
	fmt.Println(color.White("white"))
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Orange("orange"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.Gray("gray"))

	color.Style(color.Underlined)
	fmt.Println(color.White("white"))
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Orange("orange"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.Gray("gray"))
}
