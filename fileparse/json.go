package fileparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Baladiya belonging to Dairas
type Baladiya struct {
	Arabic string  `json:"arabic"`
	French string  `json:"french"`
	Lat    float64 `json:"latitude"`
	Lng    float64 `json:"longitude"`
}

// Daira belongs to Wilaya
type Daira struct {
	Arabic    string     `json:"french"`
	French    string     `json:"arabic"`
	Baladiyas []Baladiya `json:"Baladiyas"`
	Lat       float64    `json:"latitude"`
	Lng       float64    `json:"longitude"`
}

// Wilaya root parent
type Wilaya struct {
	Arabic     string   `json:"arabic"`
	French     string   `json:"french"`
	Matricule  string   `json:"matricule"`
	PhoneCodes []string `json:"phoneCodes"`
	Dairas     []Daira  `json:"Dairas"`
	Lat        float64  `json:"latitude"`
	Lng        float64  `json:"longitude"`
}

// ParseWilayaFile to parse json?
func ParseWilayaFile() (*[]Wilaya, error) {
	var wilayas []Wilaya

	fmt.Println("parsing file")

	data, err := ioutil.ReadFile("./data/Algeria.json")
	if err != nil {
		log.Fatalf("Failed to get file %s", err)
	}

	json.Unmarshal(data, &wilayas)
	return &wilayas, nil

}
