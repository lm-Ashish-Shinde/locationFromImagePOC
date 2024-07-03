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

	fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)

	address, err := reverseGeocode(latitude, longitude)
	if err != nil {
		log.Fatalf("Error reverse geocoding: %v", err)
	}

	fmt.Printf("Address: %s\n", address)

}
