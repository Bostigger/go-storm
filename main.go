package main

import (
	"bufio"
	"fmt"
	"github.com/bostigger/go-weather-tracker/controller"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var apiKey = os.Getenv("APIKEY")

	fmt.Println("Go Weather tracker")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter the city name (or enter exit to quit)")

		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				log.Fatalf("Error scanning user input: %v", err)
			}
			break
		}

		userInput := strings.Trim(scanner.Text(), " ")
		if userInput == "" {
			fmt.Println("You didn't enter any name")
			continue
		}
		if userInput == "exit" {
			fmt.Println("Exiting the program...")
			break
		}

		data, err := controller.GetWeatherData(userInput, apiKey)
		if err != nil {
			fmt.Printf("Error getting weather data: %v\n", err)
			continue
		}

		controller.ShowWeatherData(&data)
	}
}
