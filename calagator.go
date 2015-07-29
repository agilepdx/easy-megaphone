package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func sendToCalagator(eventEntry event) {
	log.Println("Getting read to send to calagator.")
	bleh, _ := json.Marshal(eventEntry)
	log.Println(string(bleh))

	token := getCalagatorAuthToken()

	sendEventToCalagator(eventEntry, token)
}

// This needs to be refactored to use something like goquery, this is p. bad
func getCalagatorAuthToken() (token string) {
	// need to scrape the new events page to make a valid request:
	// http://calagator.org/events
	//
	resp, err := http.Get("http://calagator.org/events/new")
	if err != nil {
		log.Fatalln("Couldn't read calagator web site")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// <input type="hidden" name="authenticity_token" value="HryLkfPlZ7gS7aR+x3QYSu7/l62JEtTLYs/mdVwrXY2B/qJUaVO3BVnk1QNGyv5vuKvEzPNnnWtOb8PL1Auvig==">
	re := regexp.MustCompile("<.*authenticity_token.*>")
	for _, value := range strings.Split(string(body), "\n") {
		found := re.FindStringSubmatch(value)
		if len(found) > 0 {
			// log.Println("found is ", found)
			// We now have a single line with the auth token and some utf thingy in it
			for _, innerValue := range strings.Split(found[0], "><") {
				// log.Println("Found innerValue of ", innerValue)

				if strings.Contains(innerValue, "authenticity_token") {
					newRegex := regexp.MustCompile("value=.*")
					newFound := newRegex.FindStringSubmatch(innerValue)
					if len(newFound) > 0 {
						// log.Println("newfound is ", newFound[0])
						slicedString := strings.Replace(newFound[0], "value", "", -1)

						slicedString = strings.Replace(slicedString, "\" />", "", -1)
						slicedString = strings.Replace(slicedString, "=\"", "", -1)

						// log.Println("slicedString is ", slicedString)
						return slicedString
					}
				}

			}
		}
	}

	return ""
}

func sendEventToCalagator(eventEntry event, token string) {
	// To do a test event, use the "preview" input.
	log.Println("Totally sending to calagator with auth token ", token)
}
