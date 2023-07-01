## GoStorm CLI Application

GoStorm is a Command Line Interface (CLI) application written in Go. It retrieves weather information for a specified city using OpenWeatherMap's API.

https://github.com/Bostigger/go-storm/assets/52701136/3c5e60f9-9322-47b9-8d22-3222a09bdfbf


Features

    Simple, intuitive, command-line interface.
    Fetches and displays the current weather conditions of a city.
    Provides information about the city's current temperature, minimum and maximum temperature, and atmospheric pressure.

Requirements

    Go 1.15 or higher
    An OpenWeatherMap API key (free tier available)

Setup Instructions

    Clone this repository to your local machine.

```
git clone https://github.com/bostigger/go-storm.git
```

Navigate to the project directory.

```
cd go-weather-tracker
```

Create a .env file in the root directory and add your OpenWeatherMap API key like this:

```
    APIKEY=your_api_key
```

    Replace your_api_key with your actual API key from OpenWeatherMap.

Running the Application

From the root directory of the project, run the command:

```
go run main.go
```

The program will then prompt you to enter a city name.
Using the Application
```
Enter the city name (or enter exit to quit)
Accra
```


The application will then display the current weather conditions for the city.

Output
```
Country: GH
City Name: Accra
Minimum Temperature: 26.37°C
Maximum Temperature: 26.37°C
Current Temperature: 26.37°C
Current Pressure: 1011Pa

```
To exit the application, type exit at the prompt and press enter.

Contributing

Contributions are welcome. Feel free to open a pull request or branch from this project.
