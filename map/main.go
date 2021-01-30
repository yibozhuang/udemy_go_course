package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "$FF00000",
		"green": "#75FG09",
		"white": "#600000",
	}

	// var colors map[string]string

	/*
		colors := make(map[string]string)

		colors["red"] = "#880000"
		delete(colors, "red")
	*/
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
