package seeder

// APIKey for google maps
import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var fileLocation = "./data/creds.json"

// APICredentials used to store key credentials
type APICredentials struct {
	GeoLocationKey string `json="geolocationKey"`
}

// GetAPIKey to get credential keys
func GetAPIKey() *APICredentials {
	var credentials APICredentials
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatalf("failed to find credential file", err)
		panic("You must create your own json credentials with {\"geolocationKey\": \"your google key here\"}")
	}

	json.Unmarshal(data, &credentials)
	return &credentials
}
