package fileparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Location
type Location struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
	// updateLocation(long float64, lat float64)
}

// NonZero checks for non zero values and then returns the default
func NonZero(original float64, number float64) float64 {
	if number != 0 {
		return number
	}
	return original
}

// UpdateLocation
func (loc *Location) UpdateLocation(long float64, lat float64) {
	if loc.Lng == 0 {
		loc.Lng = long
	}
	if loc.Lat == 0 {
		loc.Lat = lat
	}
}

// Baladiya belonging to Dairas
type Baladiya struct {
	Arabic string `json:"arabic"`
	French string `json:"french"`
	// Lat    float64 `json:"latitude"`
	// Lng    float64 `json:"longitude"`
	Location
}

// Daira belongs to Wilaya
type Daira struct {
	Arabic    string     `json:"french"`
	French    string     `json:"arabic"`
	Baladiyas []Baladiya `json:"Baladiyas"`
	// Lat       float64    `json:"latitude"`
	// Lng       float64    `json:"longitude"`
	Location
}

// Wilaya root parent
type Wilaya struct {
	Arabic     string   `json:"arabic"`
	French     string   `json:"french"`
	Matricule  string   `json:"matricule"`
	PhoneCodes []string `json:"phoneCodes"`
	Dairas     []Daira  `json:"Dairas"`
	// Lat        float64  `json:"latitude"`
	// Lng        float64  `json:"longitude"`
	Location
}

// ParseWilayaFile to parse json?
func ParseWilayaFile(fileLocation string) (*[]Wilaya, error) {
	var wilayas []Wilaya
	fmt.Println("parsing file")
	if fileLocation == "" {
		fileLocation = "./data/Algeria.json"
	}

	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatalf("Failed to get file %s", err)
	}

	json.Unmarshal(data, &wilayas)
	return &wilayas, nil
}
