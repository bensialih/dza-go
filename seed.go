package main

import (
	"fmt"
	seeder "geo_google/seeder"
)

func main() {
	creds := seeder.GetAPIKey()
	seeder.AddLongLat(creds.GeoLocationKey)
	fmt.Println("done")
}
