package main

import (
	// "encoding/json"
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "net/http"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	// "github.com/rwcarlsen/goexif/tiff"
	// "googlemaps.github.io/maps"
)

// Function to extract GPS coordinates from image file
func extractCoordinatesFromImage(imagePath string) (float64, float64, error) {
	// Open image file
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode EXIF data
	x, err := exif.Decode(file)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode exif data: %v", err)
	}

	// Extract GPS coordinates
	lat, long, err := x.LatLong()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to extract GPS coordinates: %v", err)
	}

	return lat, long, nil
}

// const (
// 	googleMapsAPIKey = "AIzaSyC2qJatOdSsOK1XY_TcedFOkABy297DOVw"
// )

// type GeocodingResponse struct {
// 	Results []GeocodingResult `json:"results"`
// 	Status  string            `json:"status"`
// }

// type GeocodingResult struct {
// 	FormattedAddress  string             `json:"formatted_address"`
// 	AddressComponents []AddressComponent `json:"address_components"`
// }

// type AddressComponent struct {
// 	LongName  string   `json:"long_name"`
// 	ShortName string   `json:"short_name"`
// 	Types     []string `json:"types"`
// }

// func reverseGeocode(latitude, longitude float64) (string, string, error) {
// 	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=%s", latitude, longitude, googleMapsAPIKey)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", "", fmt.Errorf("HTTP status code %d", resp.StatusCode)
// 	}

// 	var geocodingResponse GeocodingResponse
// 	err = json.NewDecoder(resp.Body).Decode(&geocodingResponse)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	if len(geocodingResponse.Results) > 0 {
// 		formattedAddress := geocodingResponse.Results[0].FormattedAddress
// 		pincode := extractPincode(geocodingResponse.Results[0])
// 		return formattedAddress, pincode, nil
// 	}

// 	return "", "", fmt.Errorf("no results found")
// }

//	func extractPincode(result GeocodingResult) string {
//		for _, addrComponent := range result.AddressComponents {
//			for _, addrType := range addrComponent.Types {
//				if addrType == "postal_code" {
//					return addrComponent.LongName
//				}
//			}
//		}
//		return ""
//	}

func reverseGeocode(latitude, longitude float64) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("failed to decode JSON: %v", err)
	}

	address := data["display_name"].(string)
	return address, nil
}

func main() {
	imagePath := "./sample.jpg"

	latitude, longitude, err := extractCoordinatesFromImage(imagePath)
	if err != nil {
		log.Fatalf("Error extracting coordinates: %v", err)
	}

	fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)

	// Reverse geocode to get address and pincode
	// address, pincode, err := reverseGeocode(latitude, longitude)
	// if err != nil {
	// 	log.Fatalf("Error reverse geocoding: %v", err)
	// }

	fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
	// fmt.Printf("Address: %s\n", address)
	// fmt.Printf("Pincode: %s\n", pincode)

	// apiKey := "YOUR_API_KEY" // Replace with your Google Maps API key
	// apiKey := "AIzaSyC2qJatOdSsOK1XY_TcedFOkABy297DOVw"
	// apiKey := "AIzaSyDUnuzfv6Vib7TvH-CZUm0MjW0_LL-CjoY"
	// apiKey := ""
	// client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	// maps.
	// if err != nil {
	// 	log.Fatalf("Failed to create Google Maps client: %v", err)
	// }

	// // Create a LatLng object
	// latlng := &maps.LatLng{
	// 	Lat: latitude,
	// 	Lng: longitude,
	// }

	// // Perform reverse geocoding
	// reverseGeocodeRequest := &maps.GeocodingRequest{
	// 	LatLng: latlng,
	// }

	// // Call the Geocode API
	// reverseGeocodeResponse, err := client.ReverseGeocode(context.Background(), reverseGeocodeRequest)
	// if err != nil {
	// 	log.Fatalf("Failed to reverse geocode: %v", err)
	// }

	// // Print the formatted address
	// if len(reverseGeocodeResponse) > 0 {
	// 	fmt.Println("Formatted Address:", reverseGeocodeResponse[0].FormattedAddress)
	// } else {
	// 	fmt.Println("No results found")
	// }

	address, err := reverseGeocode(latitude, longitude)
	if err != nil {
		log.Fatalf("Error reverse geocoding: %v", err)
	}

	fmt.Printf("Address: %s\n", address)

}
