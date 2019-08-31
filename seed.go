package main

import (
	seeder "dza-go/seeder"
	"fmt"
)

func main() {
	seeder.AddLongLat(seeder.GetAPIKey())
	fmt.Println("done")
}
