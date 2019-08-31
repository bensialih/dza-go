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

func queryGoogle(channel chan []float64, search string, client *maps.Client) {
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

	channel <- []float64{long, latitude}
}

func queryGoogleAPI(searchFor string, client *maps.Client) chan []float64 {
	channel := make(chan []float64)
	go queryGoogle(channel, searchFor, client)
	return channel
}

func seedBaladiyaLocation(baladiyas []parser.Baladiya, client *maps.Client) {
	for index, baladiya := range baladiyas {
		results := <-queryGoogleAPI(fmt.Sprintf("%s, DZ", baladiya.French), client)
		baladiyas[index].Lng = results[0]
		baladiyas[index].Lat = results[1]
		fmt.Println("baladiya long lat :> ", results)
	}
}

func seedDairaLocation(dairas []parser.Daira, client *maps.Client) {
	for index, daira := range dairas {
		results := <-queryGoogleAPI(fmt.Sprintf("%s, DZ", daira.French), client)
		dairas[index].Lng = results[0]
		dairas[index].Lat = results[1]
		fmt.Println("daira, long lat :> ", results)
		if len(daira.Baladiyas) > 0 {
			seedBaladiyaLocation(dairas[index].Baladiyas, client)
		}
	}
}

func seedWilayaLocation(wilayas []parser.Wilaya, client *maps.Client) {
	for index, wilaya := range wilayas {
		results := <-queryGoogleAPI(fmt.Sprintf("%s, DZ", wilaya.French), client)
		fmt.Println("long lat :> ", results)
		wilayas[index].Lng = results[0]
		wilayas[index].Lat = results[1]
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
