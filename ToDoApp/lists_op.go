package main

import (
	"encoding/json"
	"os"
)

type TodoList struct {
	Name string
	Description string
	Items []ListItem
}

type ListItem struct {
	Description string
	IsDone bool
}

func (tdl TodoList) Save(filename string) error {
	lStr, err := json.Marshal(tdl)
	if err != nil {
		return err
	} 
	
	// TODO: Security
	err = os.WriteFile("./data/"+filename, lStr, 0600)
	return err
}
