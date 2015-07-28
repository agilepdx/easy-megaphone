package main

import (
	"encoding/json"
	"io/ioutil"
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

	// TODO: take in a flag with a file to post.
	eventJSON, _ := readFileContents("sample-event.json")

	eventEntry := readFromJSON(eventJSON)

	// These can be refactored into a single function that calls them all
	sendToCalagator(eventEntry)
	sendToMeetup(eventEntry)
	sendToAgilePDXWebsite(eventEntry)
	sendToTwitter(eventEntry)
}

func setup() {
	// envconfig bits for various integration, such as Meetups API
}

func readFileContents(file string) ([]byte, error) {
	fileAsBytes, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalln("Couldn't open ", file)
	}

	return fileAsBytes, nil
}

func readFromJSON(eventJSON []byte) event {
	var eventEntry event
	err := json.Unmarshal([]byte(eventJSON), &eventEntry)

	if err != nil {
		log.Fatal("Couldn't parse JSON file.")
	}
	return eventEntry
}

func sendToCalagator(eventEntry event) {
	log.Println("Totally sending to calagator...")
	bleh, _ := json.Marshal(eventEntry)
	log.Println(string(bleh))
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