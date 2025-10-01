package tdl

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	var cases = [] struct{
		in string
		main rune
		expected []Parameter
	}{
		{"Create -Name=\"The list\" -Descripton=\"I'm here\"", 'c', []Parameter{{'n',"The list"}, {'d', "I'm here"}}},
		{"Update -Name=\"Todo List One!\" -Description=\"Need to do this\"", 'u', []Parameter{{'n',"Todo List One!"}, {'d', "Need to do this"}}},
	}
	
	for ti, tt := range cases {
		cmd, err := ParseInput(tt.in)
		if err != nil {
			t.Errorf("Received error when parsing %d: %v", ti, err)
		}
		validateCommand(cmd, tt.main, tt.expected, t) 
	}
}

func validateCommand(cmd Command, m rune, pars []Parameter, t *testing.T) {
	if cmd.Main != m  {
		t.Errorf("Expected %c, got %c", m, cmd.Main)
	}

	if len(cmd.Parameters) != len(pars) {
		t.Errorf("Expected %d parameters, got %d", len(pars), len(cmd.Parameters))
		fmt.Printf("Got: %v\n", cmd.Parameters)
	}

	for i, p := range pars {
		if cmd.Parameters[i] != (p){
			t.Errorf("Expected %v, got %v", p, cmd.Parameters[i])
		}
	}
}