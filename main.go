package main

import (
	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var higreen = color.New(color.FgHiGreen).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()
var hiblue = color.New(color.FgHiBlue).SprintFunc()
var hired = color.New(color.FgHiRed).SprintFunc()
var himagenta = color.New(color.FgHiMagenta).SprintFunc()
var hiwhite = color.New(color.FgHiWhite).SprintFunc()
var hiyellow = color.New(color.FgHiYellow).SprintFunc()

func main() {
	displayBanner()
	//client := newServer()
	//client.askDetails()
}
