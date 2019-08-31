package seeder

// APIKey for google maps
import (
	"fmt"
	"github.com/spf13/viper"
)

var fileLocation = "./data/creds.json"

// APICredentials used to store key credentials
type APICredentials struct {
	GeoLocationKey string `json:"geolocationKey"`
}

// GetAPIKey to get credential keys as string
func GetAPIKey() string {
	viper.SetConfigName("creds")
	viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	credentials := viper.GetString("geolocationKey")
	return credentials
}
