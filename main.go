package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
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
	GitHubToken    string
	MeetupClientID string
}

var (
	productionMode bool
	s              specification
)

// We may need to stuff an HTTP server in here to listen for Meetup.com's OAuth2 response.
// Might be good to provide a web interface for enabling new events.
func main() {
	log.Println("Starting up easy-megaphone")

	setup()

	// TODO: take in a flag with a file to post.
	eventEntry, err := fileToEvent("sample-event.json")
	if err != nil {
		log.Fatalln("Couldn't open file for reading.")
	}

	// These can be refactored into a single function that calls them all
	// sendToCalagator(eventEntry)
	sendToMeetup(eventEntry)
	// sendToAgilePDXWebsite(eventEntry)
	// sendToTwitter(eventEntry)
}

func setup() {
	// envconfig bits for various integration, such as Meetups API token
	productionMode = false

	err := envconfig.Process("easymegaphone", &s)
	if err != nil {
		log.Fatal("boo: ", err.Error())
	}

	log.Println("github token is " + s.GitHubToken)
	log.Println("meetup client id is " + s.MeetupClientID)

	if productionMode {
		log.Println("We're in production mode, gonna send real events.")
	} else {
		log.Println("We're in DEBUG mode, not sending real requests.")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/v1/meetup/", meetupReturnHandler)

	log.Println("Listening on port 8080.")
	go log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendToTwitter(eventEntry event) {
	log.Println("Totally sending to twitter...")
}
