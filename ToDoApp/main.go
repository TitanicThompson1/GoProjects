package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"time"
)

var lg *log.Logger

func init() {
	now := time.Now()
	logfn := fmt.Sprintf("logs/%d_%d_%d.txt", now.Year(), now.Month(), now.Day())
	logF, err := os.OpenFile(logfn, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Received err %v when trying to open log file", err)
	}
	lg = log.New(logF, "", log.Lmicroseconds | log.Ltime) 
}

func main() {
	lg.Println("Starting application")

	printMainMenu()
	reader := bufio.NewReader(os.Stdin)

Exit:
	for ; ; {
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			lg.Printf("Couldn't read string from stdin: %v", err)
			break Exit
		}
		text = strings.TrimSpace(text)
		lg.Println("Received: ", text)

		cmd, err := ParseInput(text)
		if err != nil {
			fmt.Println("Something went wrong while reading your command! Please try again.")
			lg.Printf("Error while parsing the input %s: %v\n", text, err)
			continue
		}

		switch cmd.Main {
			case 'Q':
				fmt.Println("Quitting application.")
				break Exit
			case 'C':
				
		}
	}
}

func printMainMenu() {
	fmt.Println(`
Welcome to the your console line ToDo app!
[S]how ToDo list
	[A]ll your lists
	[N]ame of the list you want to see
[C]reate ToDo list
	[U]ser name
	[N]ame of the list
[Q]uit application.`)

}
