package main

import (
	// "encoding/json"
	"log"
)

// TODO: whole twitter post?
type event struct {
	EventName     string `json:"EventName"`
	VenueName     string `json:"VenueName"`
	VenueLocation string `json:"VenueLocation"`
	StartTime     string `json:"StartTime"`
	EndTime       string `json:"EndTime"`
	Description   string `json:"Description"`
}

func main() {
	log.Println("Starting up easy-megaphone")

	setup()

	readFromJSON()

	sendToCalagator()
}

func setup() {
	// envconfig bits for various integration, such as Meetups API
}

func readFromJSON() {
	// read from JSON input file
}

func sendToCalagator() {
	log.Println("Sending to calagator...")
	// Totally sends to calagator
}
