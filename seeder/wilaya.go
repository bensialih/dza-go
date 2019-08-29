package seeder

import (
	"context"
	"encoding/json"
	"fmt"
	parser "geo_google/fileparse"
	"googlemaps.github.io/maps"
	"io/ioutil"
	"log"
)

var long float64
var latitude float64

func queryGoogle(search string, client *maps.Client) *[]maps.GeocodingResult {
	response, err := client.Geocode(context.Background(), &maps.GeocodingRequest{Address: search})
	if err != nil {
		log.Fatalf("error querying google maps %s", err)
	}
	return &response
}

func seedBaladiyaLocation(baladiyas []parser.Baladiya, client *maps.Client) {
	for index, baladiya := range baladiyas {
		response := *queryGoogle(fmt.Sprintf("%s, DZ", baladiya.French), client)

		if len(response) > 1 {
			fmt.Sprintln("multiple results found for %s", response[0].FormattedAddress)
		} else if len(response) == 1 {
			long = response[0].Geometry.Location.Lng
			latitude = response[0].Geometry.Location.Lat
			baladiyas[index].Lng = long
			baladiyas[index].Lat = latitude
		}
	}
}

func seedDairaLocation(dairas []parser.Daira, client *maps.Client) {
	for index, daira := range dairas {
		response := *queryGoogle(fmt.Sprintf("%s, DZ", daira.French), client)

		if len(response) > 1 {
			fmt.Sprintln("multiple results found for %s", response[0].FormattedAddress)
		} else if len(response) == 1 {
			long = response[0].Geometry.Location.Lng
			latitude = response[0].Geometry.Location.Lat

			dairas[index].Lng = long
			dairas[index].Lat = latitude

			if len(daira.Baladiyas) > 0 {
				seedBaladiyaLocation(dairas[index].Baladiyas, client)
			}
		}
	}
}

func seedWilayaLocation(wilayas []parser.Wilaya, client *maps.Client) {
	for index, wilaya := range wilayas {
		response := *queryGoogle(fmt.Sprintf("%s, DZ", wilaya.French), client)

		if len(response) > 1 {
			fmt.Sprintln("multiple results found for %s", response[0].FormattedAddress)
		} else if len(response) == 1 {
			fmt.Println("updating long/lat", long, latitude)
			long = response[0].Geometry.Location.Lng
			latitude = response[0].Geometry.Location.Lat
			wilayas[index].Lng = long
			wilayas[index].Lat = latitude

			if len(wilaya.Dairas) > 0 {
				seedDairaLocation(wilayas[index].Dairas, client)
			}
		}
	}
}

// AddLongLat adding longitude and latitude function
func AddLongLat(apiKey string) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	wilayas, err := parser.ParseWilayaFile()
	if err != nil {
		log.Fatalf("failed to parse json %s", err)
	}

	seedWilayaLocation(*wilayas, client)

	dataset, err := json.MarshalIndent(*wilayas, "", "\t")
	if err != nil {
		fmt.Sprintf("Failed to marshal wilaya object", err)
	}

	error := ioutil.WriteFile("./data/new_algeria.json", dataset, 0777)
	if error != nil {
		fmt.Println("file error", error)
	}
}
