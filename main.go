package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//s := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	var liner string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		liner = scanner.Text()
		s := cleanInput(liner)
		fmt.Printf("Your command was: %s\n", s[0])

	}

}

func cleanInput(intext string) []string {
	//s := []string{}
	lowText := strings.ToLower(intext)
	trimText := strings.TrimSpace(lowText)
	s := strings.Fields(trimText)
	return s
}
