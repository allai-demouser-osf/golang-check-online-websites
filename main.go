package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const qtdMonitoramentos = 3
const delayInSeconds = 5

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func testSite(site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error: ", err)
	} else if res.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded with success!")
	} else {
		fmt.Println("Site:", site, "was loaded with error:", res.StatusCode)
	}
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br", "https://www.google.com.br"}

	for i := 0; i < qtdMonitoramentos; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("Waiting", delayInSeconds, "seconds...")
		time.Sleep(delayInSeconds * time.Second)
	}
}

func main() {
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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
}
