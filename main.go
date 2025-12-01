package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		
		if !scanner.Scan() {
			fmt.Println("\nGoodbye!")
			return
		}

		input := scanner.Text()
		words := cleanInput(input)


		if len(words) == 0 {
			continue
		}

		firstWord := words[0]
		fmt.Println("Your command was:", firstWord)
	}
}