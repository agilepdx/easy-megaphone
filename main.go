package main

import (
	"log"
)

// TODO: whole twitter post as a field?
type event struct {
	EventName     string `json:"EventName"`
	VenueName     string `json:"VenueName"`
	VenueLocation string `json:"VenueLocation"`
	VenueDetails  string `json:"VenueDetails"`
	StartTime     string `json:"StartTime"`
	EndTime       string `json:"EndTime"`
	Description   string `json:"Description"`
}

var productionMode bool

func main() {
	log.Println("Starting up easy-megaphone")

	setup()

	// TODO: take in a flag with a file to post.
	eventEntry, err := fileToEvent("sample-event.json")
	if err != nil {
		log.Fatalln("Couldn't open file for reading.")
	}

	// These can be refactored into a single function that calls them all
	sendToCalagator(eventEntry)
	sendToMeetup(eventEntry)
	sendToAgilePDXWebsite(eventEntry)
	sendToTwitter(eventEntry)
}

func setup() {
	// envconfig bits for various integration, such as Meetups API token
	productionMode = false
}

func sendToMeetup(eventEntry event) {
	log.Println("Totally sending to meetup...")
}

func sendToAgilePDXWebsite(eventEntry event) {
	log.Println("Totally sending to agilepdx website...")
}

func sendToTwitter(eventEntry event) {
	log.Println("Totally sending to twitter...")
}
