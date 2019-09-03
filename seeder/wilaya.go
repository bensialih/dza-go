package seeder

import (
	"context"
	parser "dza-go/fileparse"
	"encoding/json"
	"fmt"
	"googlemaps.github.io/maps"
	"io/ioutil"
	"log"
)

func queryGoogle(channel *chan []float64, search string, client *maps.Client) {
	response, err := client.Geocode(context.Background(), &maps.GeocodingRequest{Address: search})
	if err != nil {
		log.Fatalf("error querying google maps %s", err)
	}
	long := 0.0
	latitude := 0.0

	if len(response) > 0 {
		long = response[0].Geometry.Location.Lng
		latitude = response[0].Geometry.Location.Lat
	}

	*channel <- []float64{long, latitude}
}

func seedBaladiyaLocation(baladiyas []parser.Baladiya, client *maps.Client) {
	for index, baladiya := range baladiyas {
		channel := make(chan []float64)
		go queryGoogle(&channel, fmt.Sprintf("%s, DZ", baladiya.French), client)
		results := <-channel
		baladiyas[index].Lng = NonZero(baladiyas[index].Lng, results[0])
		baladiyas[index].Lat = NonZero(baladiyas[index].Lat, results[1])
		fmt.Println("baladiya long lat :> ", results)
	}
}

func seedDairaLocation(dairas []parser.Daira, client *maps.Client) {
	for index, daira := range dairas {
		channel := make(chan []float64)
		go queryGoogle(&channel, fmt.Sprintf("%s, DZ", daira.French), client)
		results := <-channel
		dairas[index].Lng = NonZero(dairas[index].Lng, results[0])
		dairas[index].Lat = NonZero(dairas[index].Lat, results[1])
		fmt.Println("daira, long lat :> ", results)
		if len(daira.Baladiyas) > 0 {
			seedBaladiyaLocation(dairas[index].Baladiyas, client)
		}
	}
}

// NonZero checks for non zero values and then returns the default
func NonZero(original float64, number float64) float64 {
	if number != 0 {
		return number
	}
	return original
}

func seedWilayaLocation(wilayas []parser.Wilaya, client *maps.Client) {
	for index, wilaya := range wilayas {
		channel := make(chan []float64)
		go queryGoogle(&channel, fmt.Sprintf("%s, DZ", wilaya.French), client)
		results := <-channel
		wilayas[index].Lng = NonZero(wilayas[index].Lng, results[0])
		wilayas[index].Lat = NonZero(wilayas[index].Lat, results[1])
		if len(wilaya.Dairas) > 0 {
			seedDairaLocation(wilayas[index].Dairas, client)
		}
	}

}

// AddLongLat adding longitude and latitude function
func AddLongLat(apiKey string) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	fileLocation := "./data/new_algeria.json"
	wilayas, err := parser.ParseWilayaFile(fileLocation)
	if err != nil {
		log.Fatalf("failed to parse json %s", err)
	}

	seedWilayaLocation(*wilayas, client)

	dataset, err := json.MarshalIndent(*wilayas, "", "\t")
	if err != nil {
		fmt.Printf("Failed to marshal wilaya object %s", err)
	}

	error := ioutil.WriteFile("./data/new_algeria.json", dataset, 0777)
	if error != nil {
		fmt.Println("file error", error)
	}
}
