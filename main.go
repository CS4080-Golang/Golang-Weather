package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var httpClient http.Client

func main() {
	// Provide a timeout of 10 secs to avoid waiting indefinitely
	httpClient = http.Client{
		Timeout: time.Second * 10,
	}

	// standard input and output
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nPlease provide the location: ")
	scanner.Scan()
	inputAddress := scanner.Text()
	fmt.Print("Which unit would you prefer (Metric or Imperial)? (default is metric): ")
	scanner.Scan()
	weatherUnit := scanner.Text()
	fmt.Print("Do you want to check the current weather, hourly forecast, or daily forecast? (default is current): ")
	scanner.Scan()
	weatherPeriod := scanner.Text()

	// process the unit input
	strings.ToLower(weatherUnit)
	if strings.Contains(weatherUnit, "met") {
		weatherUnit = MetricUnits
	} else if strings.Contains(weatherUnit, "imp") {
		weatherUnit = ImperialUnits
	} else {
		weatherUnit = MetricUnits
	}

	// process the period input
	strings.ToLower(weatherPeriod)
	if strings.Contains(weatherPeriod, "cur") {
		weatherPeriod = CurrentWeather
	} else if strings.Contains(weatherPeriod, "hou") {
		weatherPeriod = HourlyWeather
	} else if strings.Contains(weatherPeriod, "dai") {
		weatherPeriod = DailyWeather
	} else {
		weatherPeriod = CurrentWeather
	}

	// request Google Geocoding to convert address to latitude and longitude
	latLng, err := convertAddressToLatLng(inputAddress)
	if err != nil {
		panic(err)
	}

	// request OpenWeather API with latitude/longitude, unit, and period
	weather, err := reqWeatherDataWithLatLng(latLng, weatherUnit, weatherPeriod)
	if err != nil {
		panic(err)
	}

	// print OpenWeather response depending on the uesr selected weather period
	switch weatherPeriod {
	case CurrentWeather:
		displayWeatherData(*weather.Current, inputAddress, weatherUnit)
	case HourlyWeather:
		displayWeatherData(*weather.Hourly, inputAddress, weatherUnit)
	case DailyWeather:
		displayWeatherData(*weather.Daily, inputAddress, weatherUnit)
	}

}

// displayWeatherData loops and displays through the Weather Data returned from OpenWeather API
func displayWeatherData(weather interface{}, userAddress string, units string) {
	fmt.Printf("\nWeather data for %s\n", userAddress)

	switch weather.(type) {
	case CurrentWeatherApiResponse:
		fmt.Print(weather.(CurrentWeatherApiResponse).toString(units))
	case []HourlyWeatherApiResponse:
		for _, hourlyData := range weather.([]HourlyWeatherApiResponse) {
			fmt.Print(hourlyData.toString(units))
		}
	case []DailyWeatherApiResponse:
		for _, dailyData := range weather.([]DailyWeatherApiResponse) {
			fmt.Print(dailyData.toString(units))
		}
	}
}
