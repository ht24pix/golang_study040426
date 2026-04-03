package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Go Task Manager v0.1")
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Print(">")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
				fmt.Println("Usage: add <task description>")
				continue
			}

			tasks = append(tasks, parts[1])
			fmt.Println("Added task")
		case "list":
			if len(tasks) == 0 {
				fmt.Println("no tasks")
				continue
			}
			for i, t := range tasks {
				fmt.Printf("%d. %s\n", i+1, t)
			}
		default:
			fmt.Println("Unknown Command. Use add, list, quit.")
		}

		if line == "" {
			fmt.Println("Please enter the command.")
			continue
		}
		if line == "quit" {
			fmt.Println("Goodbye")
			break
		}
	}
}
