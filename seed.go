package main

import (
	seeder "dza-go/seeder"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("time elapsed : ", time.Since(now).Seconds())
	}()

	seeder.AddLongLat(seeder.GetAPIKey())
	fmt.Println("done")
}
