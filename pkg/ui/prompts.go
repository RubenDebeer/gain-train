package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PromptString(prompt string) string {
	fmt.Printf("%s: ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func PromptInt(prompt string) int {
	for {
		input := PromptString(prompt)
		if value, err := strconv.Atoi(input); err == nil {
			return value
		}
		fmt.Println("Please enter a valid number.")
	}
}

func PromptConfirm(prompt string) bool {
	for {
		input := strings.ToLower(PromptString(prompt + " (y/n)"))

		if input == "y" || input == "yes" {
			return true
		}

		if input == "n" || input == "no" {
			return false
		}

		fmt.Println("Please enter 'y' for yes or 'n' for no.")
	}
}

func PrintSuccess(message string) {
	fmt.Printf("Yeah Babby!!! %s\n", message)
}

func PrintError(message string) {
	fmt.Printf("Nah Bru...! %s\n", message)
}

func PrintInfo(message string) {
	fmt.Printf("FYI %s\n", message)
}
