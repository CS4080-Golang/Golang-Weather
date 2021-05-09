package main

// A LatLng object containing the latitude and longitude of the user provided address
type LatitudeLongitude struct {
	Lat float64
	Lng float64
}

// A GoogleGeocodeResult object to store the Longitute and Latitude from the result of the Google Geocode response
type GoogleGeocodeResult struct {
	Geometry struct {
		Location LatitudeLongitude
	}
}

// A GoogleGeocodeResponse object contains the response from Google Geocoding API
type GoogleGeocodeResponse struct {
	Status  string
	Results []GoogleGeocodeResult
}
