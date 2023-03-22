package main

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	// Send an HTTP GET request to the geolocation API
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse the response body using the gjson library
	body := gjson.ParseReader(resp.Body)

	// Extract the latitude and longitude from the response
	latitude := body.Get("lat").Float()
	longitude := body.Get("lon").Float()

	// Print the latitude and longitude
	fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
}
