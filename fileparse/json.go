package fileparse

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Name of the place
type Name struct {
	Arabic string `json:"arabic"`
	French string `json:"french"`
}

// Location of place
type Location struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// UpdateLocation for locations
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
	Location
	Name
}

// Daira belongs to Wilaya
type Daira struct {
	Name
	Baladiyas []Baladiya `json:"Baladiyas"`
	Location
}

// Wilaya root parent
type Wilaya struct {
	Name
	Matricule  string   `json:"matricule"`
	PhoneCodes []string `json:"phoneCodes"`
	Dairas     []Daira  `json:"Dairas"`
	Location
}

//GetFileContent of file
func GetFileContent(fileLocation string) []byte {
	fmt.Println("parsing file")
	if fileLocation == "" {
		fileLocation = "./data/Algeria.json"
	}

	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatalf("Failed to get file %s", err)
	}
	return data
}
