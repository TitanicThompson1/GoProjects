package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Command struct {
	Main rune
	Secondary []rune
}


func ParseInput(in string) (Command, error){
	parsed := strings.Split(in, " ")
	if len(parsed) <= 0 {
		return Command{}, fmt.Errorf("Error while parsing string %s", in)
	}

	cmd := Command{}
	cmd.Main = unicode.ToLower(rune(parsed[0][0]))
	cmd.Secondary = make([]rune, len(parsed) - 1)
	for i := 1; i < len(parsed); i++ {
		cmd.Secondary[i - 1] =  unicode.ToLower(rune(parsed[i][0]))
	}
	return cmd, nil
}