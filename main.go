package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(intext string) []string {
	//s := []string{}
	lowText := strings.ToLower(intext)
	trimText := strings.TrimSpace(lowText)
	s := strings.Fields(trimText)
	return s
}
