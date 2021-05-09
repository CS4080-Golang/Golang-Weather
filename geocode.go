package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Takes an address and returns an object containing its equivalent latitude and longitude
func convertAddressToLatLng(userAddress string) (latLng LatitudeLongitude, err error) {

	/* add escape characters to the address to avoid commas or spaces in the user address
	example: 1600 Amphitheatre Parkway, Mountain View, CA becomes 1600+Amphitheatre+Parkway,+Mountain+View,+CA
	*/
	escapeAddress := url.QueryEscape(userAddress)

	// format the request string according to Google API docs
	reqURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s",
		GoogleGeocodingAPIKey,
		escapeAddress,
	)

	// execute request and handle errors if any
	res, err := httpClient.Get(reqURL)
	if err != nil {
		return latLng, err
	}

	// close the response body after finishing with it (required in Go docs)
	defer res.Body.Close()

	var geocodingRes GoogleGeocodeResponse

	err = json.NewDecoder(res.Body).Decode(&geocodingRes)
	if err != nil {
		return latLng, err
	}

	// check response status and validate response
	if geocodingRes.Status != "OK" || len(geocodingRes.Results) < 1 {
		return latLng, err
	}

	return geocodingRes.Results[0].Geometry.Location, err
}
