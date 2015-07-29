package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

// Remote Address:162.242.171.114:80
// Request URL:http://calagator.org/events
// Request Method:POST
// Status Code:200 OK

// request headers
// POST /events HTTP/1.1
// Host: calagator.org
// Connection: keep-alive
// Content-Length: 384
// Cache-Control: max-age=0
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
// Origin: http://calagator.org
// Upgrade-Insecure-Requests: 1
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36
// Content-Type: application/x-www-form-urlencoded
// Referer: http://calagator.org/events/new
// Accept-Encoding: gzip, deflate
// Accept-Language: en-US,en;q=0.8
// Cookie: calagator_session=BAh7B0kiD3Nlc3Npb25faWQGOgZFVEkiJTdjY2Q3MWVhMjU3MmRlMWFhMzQ5YzhjZTRkZGY1ZmQ2BjsAVEkiEF9jc3JmX3Rva2VuBjsARkkiMXNKQitRRHIvcHhQbGxZT1lUenk4YW1aWUdhbkpYTHJFQ0NEZzlmcERCM1U9BjsARg%3D%3D--70b5026f71a656da75ba240f7354cc393e3fca8b; __utma=50912187.1139943222.1417543983.1433348881.1433785727.15; __utmc=50912187; __utmz=50912187.1431401605.9.6.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); _gat=1; _calagator_org_session=RWtMaURxVEUvVGxtdFRpMjBkODZPb0Z0RjZ5T2JyWCtqTUVzc2ZVanpPSTBBUDNabU1FRzIyRnVRZDNEQWtvaTFtbWVpcUNjdE00UzFnTVFHTTZlNkhsS3lKUk9FNFl3OXZvWXhiMjBxTTNuMDFMUXdZT1VMZ0M3ekhEelBIUk15S2lRcEw0Y1QxRGF3cDVuMEo1RWtBPT0tLWZkczBDRzFvRysrWkh0dEVYYnVTc1E9PQ%3D%3D--c41aad17688bb3cf5e0dc2b9047f4ab9cbf1f9aa; _ga=GA1.2.1139943222.1417543983

// POST form data
// utf8:✓
// authenticity_token:W6uNk9odCe5kyhzcCMrhaUUWgTRv+GcLWbnu9rO58XLE6aRWQKvZUy/DbaGJdAdME0LSVRWNLqt1GctIO5kDdQ==
// event[title]:tester
// venue_name:
// event[venue_id]:
// start_date:2015-07-26
// start_time:8:00 PM
// end_date:2015-07-26
// end_time:9:00 PM
// event[url]:
// event[description]:
// event[venue_details]:
// event[tag_list]:
// trap_field:
// preview:Preview

func makePostFormValues(eventEntry event, authToken string) (values url.Values) {
	vals := url.Values{}

	vals.Set("utf8", "✓")
	vals.Set("authenticity_token", authToken)
	vals.Set("event[title]", eventEntry.EventName)
	vals.Set("venue_name", eventEntry.VenueName)
	vals.Set("start_date", getDateFromDateTime(eventEntry.StartTime)) // 2015-07-26
	vals.Set("start_time", getTimeFromDateTime(eventEntry.StartTime)) // 8:00 PM
	vals.Set("end_date", getDateFromDateTime(eventEntry.EndTime))
	vals.Set("end_time", getTimeFromDateTime(eventEntry.EndTime))
	vals.Set("event[url]", "")
	vals.Set("event[description]", eventEntry.Description)
	vals.Set("event[venue_details]", "")
	vals.Set("event[tag_list]", "testertag")
	vals.Set("trap_field", "")
	vals.Set("preview", "Preview")

	return vals
}

func sendEventToCalagator(eventEntry event, token string) {
	// To do a test event, use the "preview" input.
	log.Println("Sending to calagator with auth token ", token)

	resp, err := http.PostForm("http://calagator.org/events", makePostFormValues(eventEntry, token))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln("Fatal error posting form: ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
	if resp.StatusCode != 200 {
		log.Fatalln("Oops got status code ", resp.Status)
	}
	log.Println("We totally posted to calagator, got a 200 back.")
}
