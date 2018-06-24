package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f74325",
		"white": "#ffffff",
	}

	printMap(colors)

	//fmt.Println(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {

		fmt.Println(key)
		fmt.Println(value)
	}
}
