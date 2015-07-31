package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type githubPullRequestResponse struct {
	RequestNumber int `json:"number"`
}

func sendToAgilePDXWebsite(eventEntry event) {
	log.Println("Totally sending to agilepdx website...")
	createCommitOnBranch(eventEntry)
}

// https://developer.github.com/v3/
func createCommitOnBranch(eventEntry event) {
	// verify we're in the right repo
	var cmdOut []byte
	var err error

	websiteDir := "/Users/matthewmayer/Documents/agilepdx/agilepdx.github.io"
	err = os.Chdir(websiteDir)

	if err != nil {
		log.Fatalln("Boo, couldn't change dir: ", err)
	}

	cmd := "git"
	args := []string{"remote", "-v"}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git: ", os.Stderr, err)
	}
	// log.Println("git remote output: ", string(cmdOut))

	if strings.Contains(string(cmdOut), "git@github.com:agilepdx/agilepdx.github.io") {
		log.Println("We're in the right spot.")
	}

	// git checkout master; git pull
	cmd = "git"
	args = []string{"checkout", "master"}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git checkout master: ", os.Stderr, err)
	}

	cmd = "git"
	args = []string{"pull"}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git pull: ", os.Stderr, err)
	}

	// git checkout -b feature-new-event-[timestamp]
	branchName := fmt.Sprintf("feature-mah-event-%v", time.Now().Unix())
	log.Println("Gonna make a branch of name ", branchName)
	cmd = "git"
	args = []string{"checkout", "-b", branchName}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git checkout -b: ", os.Stderr, err)
	}

	// update web site index.html with new event
	updateEventsListing(eventEntry)
	// but for testing:
	f, err := os.OpenFile(websiteDir+"/index.html", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Println("Couldn't open index.html: ", err)
	}
	_, err = f.WriteString("<!-- test -->")
	if err != nil {
		log.Println("Couldn't write to index.html: ", err)
	}
	f.Close()

	// git add, git commit -m
	cmd = "git"
	args = []string{"add", "index.html"}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git add: ", os.Stderr, err)
	}

	cmd = "git"
	args = []string{"commit", "-m", "\"easy-megaphone: updated upcoming events.\""}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git commit: ", os.Stderr, err)
	}

	// git push --set-upstream origin branchname
	cmd = "git"
	args = []string{"push", "--set-upstream", "origin", branchName}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git push --set-upstream: ", os.Stderr, err)
	}

	// create pull request via github API
	pullRequestPayload := []byte(fmt.Sprintf(`{"title":"easy-megaphone automated updated",
    "body" : "Automatic pull request on behalf of easy-megaphone.", "head" : "%v" , "base" : "master"}`, branchName))

	// curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/agilepdx/agilepdx.github.io/pulls", bytes.NewBuffer(pullRequestPayload))
	req.Header.Add("Authorization", "token "+s.GitHubToken)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln("Blew up asking github to make a PR.")
	}

	var gitHubPullRequest githubPullRequestResponse
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Couldn't read response json: ", err)
	}
	err = json.Unmarshal(respBody, &gitHubPullRequest)
	if err != nil {
		log.Println("Couldn't parse response json: ", err)
	}

	// accept PR via GH API
	// PUT /repos/:owner/:repo/pulls/:number/merge
	time.Sleep(100 * time.Millisecond) // Was having issues with GH, this can probably be pulled.
	pullRequestPayload = []byte(fmt.Sprintf(`{"commit_message":"easy-megaphone automated update"}`))
	pullRequestURL := "https://api.github.com/repos/agilepdx/agilepdx.github.io/pulls/" + strconv.Itoa(gitHubPullRequest.RequestNumber) + "/merge"
	req, err = http.NewRequest("PUT", pullRequestURL, bytes.NewBuffer(pullRequestPayload))
	req.Header.Add("Authorization", "token "+s.GitHubToken)
	resp, err = client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln("Blew up asking github to merge a PR.")
	}

	// delete branch via a local push with --delete
	cmd = "git"
	args = []string{"push", "--delete", "origin", branchName}
	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal("Couldn't run git push delete: ", os.Stderr, err, string(cmdOut))
	}
}

func updateEventsListing(eventEntry event) {
	// shell magicks:

	// grep to remove last event

	// sed to move event2 to event3

	// sed to move event1 to event2

	// insert eventEntry as event1
}
