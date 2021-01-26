package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("116.25.147.110")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("City name: %v\n", record.City.Names["en"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("Subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Printf("Country name: %v\n", record.Country.Names["en"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
