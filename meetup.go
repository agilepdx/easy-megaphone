package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	code             string
	meetupEventEntry event
)

func sendToMeetup(eventEntry event) {
	log.Println("Totally sending to meetup...")
	meetupEventEntry = eventEntry
}

func handler(w http.ResponseWriter, r *http.Request) {
	meetupAuthURL := fmt.Sprintf("https://secure.meetup.com/oauth2/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost:8080/v1/meetup/", s.MeetupClientID)
	fmt.Fprintf(w, "<html><body>Head on over to the <a href=\"%s\">meetup auth page</a> to authorize easy-megaphone for meetup.com</body></html>", meetupAuthURL)

	w.Header().Set("Content-Type", "text/html")
}

// This may need to handle both the code return and the json body return with the auth token in it
func meetupReturnHandler(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query()) < 1 {
		log.Fatal("Didn't get anything in the response query string from Meetup")
	}
	code = r.URL.Query()["code"][0]

	log.Println("extracted code is ", code)

	requestAccessToken()
	createEvent(meetupEventEntry) // from global var, set in sendToMeetup
}

func makePostFormValuesForMeetup(eventEntry event) (values url.Values) {
	vals := url.Values{}

	vals.Set("client_id", s.MeetupClientID)
	vals.Set("client_secret", s.MeetupClientSecret)
	vals.Set("grant_type", "authorization_code")
	vals.Set("redirect_uri", "http://localhost:8080/v1/meetup/")
	vals.Set("code", code)

	return vals
}

func requestAccessToken() {
	resp, err := http.PostForm("https://secure.meetup.com/oauth2/access", makePostFormValuesForMeetup(meetupEventEntry))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln("Fatal error posting form to meetup for auth token: ", err)
	}

	if resp.StatusCode == 200 || resp.StatusCode == 302 {
		log.Println("We posted to meetup, got ", resp.Status)
		// extract access_token from json body

	} else {
		log.Println("Didn't get a 200 or 302 back.  Got: ", resp.Status)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(string(body))
		}
	}

}

func createEvent(eventEntry event) {
	if productionMode {
		// todo: create body
		// also: content-type?
		resp, err := http.PostForm("http://www.meetup.com/meetup_api/docs/2/event/#create", nil)
		defer resp.Body.Close()
		if err != nil {
			log.Fatalln("Fatal error posting form to meetup: ", err)
		}
		// body, err := ioutil.ReadAll(resp.Body)
		// log.Print(string(body))
		if resp.StatusCode == 200 || resp.StatusCode == 302 {
			// 200 means the non-productionMode request succeeded
			// 302 means the productionMode request made a new event on calagator
			log.Println("We posted to meetup, got ", resp.Status)
		} else {
			log.Println("Didn't get a 200 or 302 back.  Got: ", resp.Status)
		}
		// create event:
		// http://www.meetup.com/meetup_api/docs/2/event/#create
		// curl -H "Authorization: Bearer {access_token}" https://api.meetup.com/2/member/self/
	}
}
