package main

import (
	"fmt"
	"time"
)

/* NOTE: Check OpenWeather API Doc for variables reference https://openweathermap.org/api/one-call-api */

const (
	CurrentWeather  = "current"
	MinutelyWeather = "minutely"
	HourlyWeather   = "hourly"
	DailyWeather    = "daily"
	ImperialUnits   = "imperial"
	MetricUnits     = "metric"
)

// Condition is embedded into Current, Hourly, Daily weather object
type Condition struct {
	Description string
}

// CurrentWeatherApiResponse is an object containing Current Moment data from OpenWeather OneCall API
type CurrentWeatherApiResponse struct {
	Dt         int64
	Temp       float32
	Humidity   int
	Visibility int
	Wind_speed float32
	Weather    []Condition
}

// toString function for the CurrentWeatherApiResponse struct
func (weather CurrentWeatherApiResponse) toString(unit string) string {
	var tempUnit string
	var SpeedUnit string

	if unit == MetricUnits {
		tempUnit = "C"
		SpeedUnit = "m/s"
	} else {
		tempUnit = "F"
		SpeedUnit = "mi/hr"
	}

	time := time.Unix(weather.Dt, 0)

	return fmt.Sprintf("Date: %2d/%d | Weekday: %-9s | Current temperature: %g%s | Humidity: %d%% | Visibility: %dm | Wind Speed: %g %s | Condition: %s\n",
		time.Month(),
		time.Day(),
		time.Weekday().String(),
		weather.Temp,
		tempUnit,
		weather.Humidity,
		weather.Visibility,
		weather.Wind_speed,
		SpeedUnit,
		weather.Weather[0].Description,
	)
}

// HourlyWeatherApiResponse is an object containing Hourly data from OpenWeather OneCall API
type HourlyWeatherApiResponse struct {
	Dt         int64
	Temp       float32
	Humidity   int
	Visibility int
	Wind_speed float32
	Weather    []Condition
}

// toString function for the HourlyWeatherApiResponse struct
func (weather HourlyWeatherApiResponse) toString(unit string) string {
	var tempUnit string
	var SpeedUnit string

	if unit == MetricUnits {
		tempUnit = "C"
		SpeedUnit = "m/s"
	} else {
		tempUnit = "F"
		SpeedUnit = "mi/hr"
	}

	time := time.Unix(weather.Dt, 0)

	return fmt.Sprintf("%-9s %2d/%2d %02d:00 | Temperature: %5.2f%s | Humidity: %d%% | Visibility: %dm | Wind Speed: %5g %s | Condition: %s\n",
		time.Weekday().String(),
		time.Month(),
		time.Day(),
		time.Hour(),
		weather.Temp,
		tempUnit,
		weather.Humidity,
		weather.Visibility,
		weather.Wind_speed,
		SpeedUnit,
		weather.Weather[0].Description,
	)
}

// DailyWeatherApiResponse is an object containing Daily data from OpenWeather OneCall API
type DailyWeatherApiResponse struct {
	Dt   int64
	Temp struct {
		Min float32
		Max float32
	}
	Humidity   int
	Visibility int
	Wind_speed float32
	Weather    []Condition
}

// toString function for the DailyWeatherApiResponse struct
func (weather DailyWeatherApiResponse) toString(unit string) string {
	var tempUnit string
	var SpeedUnit string

	if unit == MetricUnits {
		tempUnit = "C"
		SpeedUnit = "m/s"
	} else {
		tempUnit = "F"
		SpeedUnit = "mi/hr"
	}

	time := time.Unix(weather.Dt, 0)
	return fmt.Sprintf("%-9s %2d/%2d | High: %5.2f%s Low: %5.2f%s | Humidity: %d%% | Visibility: %dm | Wind Speed: %5g %s | Condition: %s\n",
		time.Weekday().String(),
		time.Month(),
		time.Day(),
		weather.Temp.Max,
		tempUnit,
		weather.Temp.Min,
		tempUnit,
		weather.Humidity,
		weather.Visibility,
		weather.Wind_speed,
		SpeedUnit,
		weather.Weather[0].Description,
	)
}

// WeatherOneCallApiResponse contains the response from OpenWeather OneCall API
type WeatherOneCallApiResponse struct {
	Current *CurrentWeatherApiResponse
	Hourly  *[]HourlyWeatherApiResponse
	Daily   *[]DailyWeatherApiResponse
}
