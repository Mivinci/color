# 🌈 color
Color helps you print decorated text to the terminal including colors, brightness, *italic* and <u>underlined</u> styles.

[![Build Status](https://travis-ci.org/Mivinci/color.svg?branch=master)](https://travis-ci.org/Mivinci/color)

## 🚀 Installation

```go
go get github.com/mivinci/color
```

## 📖 Usage

```go
fmt.Println(color.Red("text"))
```

This will print a plain red text with no other decorations.

If you need italic style text, just switch the style before getting color

```go
color.Style(color.Italic)
fmt.Println(color.Red("text"))
```

Now the text printed out will be italicized like "*text*".

## 👚 Styles

You can run [_examples/main.go](_examples/main.go)  to see all the colors and styles supported.

## 📦 Projects Using Color

|                Package                |                Author                 |                     Description                     |
| :-----------------------------------: | :-----------------------------------: | :-------------------------------------------------: |
| [nut](https://github.com/mivinci/nut) | [Mivinci](https://github.com/mivinci) | Simple, fast, easy2use web framework (experimental) |

If you find [color]() works great in your project, please remember to put your package in this table via [PR](https://github.com/Mivinci/color/pulls). Thanks! 🙃

