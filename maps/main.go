package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"black": "#000000",
		"white": "#ffffff",
		"red":   "#ff0000",
	}
	// fmt.Println(colors)

	colors["yellow"] = "#FFFF00"
	// fmt.Println(colors)

	delete(colors, "black")
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
