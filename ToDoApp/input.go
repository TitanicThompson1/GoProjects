package tdl

import (
	"unicode"
)

type Command struct {
	Main rune
	Parameters []Parameter
}

type Parameter struct {
	Name rune
	Value string
}

func ParseInput(in string) (Command, error){
	cmd := Command{}
	cmd.Main = unicode.ToLower(rune(in[0]))
	cmd.Parameters = make([]Parameter, 16)
	pi := 0
	for i := 1; i < len(in); i++ {
		switch in[i]{
			case '-':
				cmd.Parameters[pi]= parseParameter(in, &i)
				pi++
			default:
				continue
		}
	}
	cmd.Parameters = cmd.Parameters[:pi]
	return cmd, nil
}

func parseParameter(in string, i* int) Parameter{
	ti := *i + 1
	p := Parameter{Name: unicode.ToLower(rune(in[ti]))}
	
	// After this, we passed the " and are at the value
	for ; in[ti] != '"'; ti++ {
	}
	ti++

	tmp := make([]byte, 256)
	j := 0
	for ; ti < len(in) && in[ti] != '"'; ti++ {
		tmp[j] = in[ti]
		j++ 
	}
	
	p.Value = string(tmp[:j])
	*i = ti
	return p
}