package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println(command)
	return command
}

func main() {
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
		os.Exit(-1)
	}
}
