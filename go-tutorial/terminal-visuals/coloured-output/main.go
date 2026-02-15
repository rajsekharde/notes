package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
    Info    = "\033[1;34m%s\033[0m"
    Success = "\033[1;32m%s\033[0m"
    Warning = "\033[1;33m%s\033[0m"
    Error   = "\033[1;31m%s\033[0m"
    Reset   = "\033[0m" // Resets the color and style
)

func main() {

	// Using built in functions

    fmt.Printf(Info, "This is an info message.\n")
    fmt.Printf(Success, "This is a success message.\n")
    fmt.Printf(Warning, "This is a warning message.\n")
    fmt.Printf(Error, "This is an error message.\n")

    // You can also concatenate the codes directly
    fmt.Println("\033[1;36m" + "This is cyan and bold text" + Reset)

	fmt.Println()

	fmt.Printf("\033[1;31m%s\033[0m\n", "Red - 31")
	fmt.Printf("\033[1;32m%s\033[0m\n", "Green - 32")
	fmt.Printf("\033[1;33m%s\033[0m\n", "Yellow - 33")
	fmt.Printf("\033[1;34m%s\033[0m\n", "Blue - 34")
	fmt.Printf("\033[1;35m%s\033[0m\n", "Purple - 35")
	fmt.Printf("\033[1;36m%s\033[0m\n", "Cyan - 36")
	fmt.Printf("\033[1;37m%s\033[0m\n", "White - 37")

	// Using color library

	color.Blue("This is a blue text")
	color.RGB(0, 0, 255).Println("This is RGB blue text")
	color.RGB(0, 255, 0).Println("This is RGB green text")
}