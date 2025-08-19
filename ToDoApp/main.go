package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"time"
)

func main() {
	
	now := time.Now()
	logfn := fmt.Sprintf("logs/%d_%d_%d.txt", now.Year(), now.Month(), now.Day())
	logF, err := os.OpenFile(logfn, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Received err %v when trying to open log file", err)
	}
	defer logF.Close()

	lg := log.New(logF, "", log.Lmicroseconds | log.Ltime) 
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
		fmt.Println("Received: ", text)
		switch text {
			case "Q":
				println("Goodbye!")
				break Exit
		}
	}
}

func printMainMenu() {
	fmt.Println(`
Welcome to the your console line ToDo app!
[S]how ToDo list
	[A]ll your lists
	[N]ame of the list you want to see
[Q]uit application.`)

}
