package main

import (
	"log"
)

func sendToMeetup(eventEntry event) {
	log.Println("Totally sending to meetup...")
}

// Each user may need to set up an OAuth2 consumer in Meetup once, not sure...

// Tell user to visit https://secure.meetup.com/oauth2/authorize
// ?client_id=YOUR_CONSUMER_KEY
// &response_type=code
// &redirect_uri=YOUR_CONSUMER_REDIRECT_URI

// user gets redirected to localhost:8080/whatever
// extract "code" from response

// use "code" to send a request to https://secure.meetup.com/oauth2/access
// with body:
// client_id=YOUR_CONSUMER_KEY
// &client_secret=YOUR_CONSUMER_SECRET
// &grant_type=authorization_code
// &redirect_uri=SAME_REDIRECT_URI_USED_FOR_PREVIOUS_STEP
// &code=CODE_YOU_RECEIVED_FROM_THE_AUTHORIZATION_RESPONSE

// snag access token and refresh token
// "access_token"
// "refresh_token"

// create event:
// http://www.meetup.com/meetup_api/docs/2/event/#create
// curl -H "Authorization: Bearer {access_token}" https://api.meetup.com/2/member/self/
