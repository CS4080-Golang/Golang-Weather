package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func reqWeatherDataWithLatLng(latLng LatitudeLongitude, units string, period string) (weather WeatherOneCallApiResponse, err error) {
	// slice containing excluded periods (current, minutely, hourly, daily)
	exclude := []string{MinutelyWeather}

	if period != CurrentWeather {
		exclude = append(exclude, CurrentWeather)
	}
	if period != HourlyWeather {
		exclude = append(exclude, HourlyWeather)
	}
	if period != DailyWeather {
		exclude = append(exclude, DailyWeather)
	}

	excludePeriods := strings.Join(exclude, ",")

	reqURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?appid=%s&lat=%g&lon=%g&exclude=%s&units=%s",
		OpenWeatherAPIKey,
		latLng.Lat,
		latLng.Lng,
		excludePeriods,
		units,
	)

	res, err := httpClient.Get(reqURL)
	if err != nil {
		return weather, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return weather, fmt.Errorf("OpenWeather Request Failed: %s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&weather)

	return weather, err
}
