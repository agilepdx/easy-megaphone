package main

import (
	"github.com/kelseyhightower/envconfig"
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

type specification struct {
	GitHubToken string
}

var (
	productionMode bool
	s              specification
)

func main() {
	log.Println("Starting up easy-megaphone")

	setup()
	if productionMode {
		log.Println("We're in production mode, gonna send real events.")
	} else {
		log.Println("We're in DEBUG mode, not sending real requests.")
	}
	// TODO: take in a flag with a file to post.
	eventEntry, err := fileToEvent("sample-event.json")
	if err != nil {
		log.Fatalln("Couldn't open file for reading.")
	}

	// These can be refactored into a single function that calls them all
	//sendToCalagator(eventEntry)
	sendToMeetup(eventEntry)
	sendToAgilePDXWebsite(eventEntry)
	sendToTwitter(eventEntry)
}

func setup() {
	// envconfig bits for various integration, such as Meetups API token
	productionMode = false

	err := envconfig.Process("easymegaphone", &s)
	if err != nil {
		log.Fatal("boo: ", err.Error())
	}

	log.Println("github token is " + s.GitHubToken)
}

func sendToMeetup(eventEntry event) {
	log.Println("Totally sending to meetup...")
}

func sendToTwitter(eventEntry event) {
	log.Println("Totally sending to twitter...")
}
