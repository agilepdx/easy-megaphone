package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func sendToAgilePDXWebsite(eventEntry event) {
	log.Println("Totally sending to agilepdx website...")
	createCommitOnBranch()
}

// https://developer.github.com/v3/
func createCommitOnBranch() {
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
	log.Println("git output: ", string(cmdOut))
	// TODO: change this to agilepdx/agilepdx.github.io
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
	f, err := os.OpenFile(websiteDir+"/index.html", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Println("Couldn't open index.html: ", err)
	}
	_, err = f.WriteString("FOOO")
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
	// ~magick~

	// accept PR via GH API

	// delete branch via GH API
}
