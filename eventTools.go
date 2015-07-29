package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func fileToEvent(fileName string) (event, error) {
	eventJSON, _ := readFileContents(fileName)
	eventEntry := readFromJSON(eventJSON)

	err := vetEvent(eventEntry)
	if err != nil {
		log.Fatalln("Issues with JSON file: ", err)
	}

	return eventEntry, nil
}

// vetEvent vets the event read in for basic validity
func vetEvent(eventEntry event) error {
	// good event:
	// {
	//  "EventName" : "Test event",
	// 	"VenueName" : "Test venue",
	// 	"VenueLocation" : "Test venue location",
	//  "VenueDetails" : "Test venue details",
	// 	"StartTime" : "2015-07-26|8:00 PM",
	// 	"EndTime" : "2015-07-26|9:00 PM",
	// 	"Description" : "Looong event description, trust me"
	// }

	return nil
}

func getDateFromDateTime(dateTime string) string {
	date := strings.Split(dateTime, "|")[0]
	return date
}

func getTimeFromDateTime(dateTime string) string {
	time := strings.Split(dateTime, "|")[1]
	return time
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
