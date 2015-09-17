# easy-megaphone
Take in JSON, talk to publication services

## Outline

1. Take in JSON file (format TBD).
2. Output to various services

Services should include:
* Meetup.com API (in progress)
* Twitter API
* Calagator (done)
* AgilePDX website via GitHub Pages (done)

## Credentials

Store in environment.

Github integration relies on a [personal access token](https://developer.github.com/v3/oauth/).

Meetup uses OAuth2 and needs a client ID and client secret.

`cat ~/easy-megaphone-creds.sh`

```
#!/bin/sh

export EASYMEGAPHONE_GITHUBTOKEN="token here"
export EASYMEGAPHONE_MEETUPCLIENTSECRET="client secret"
export EASYMEGAPHONE_MEETUPCLIENTID="client id"
```

### Integrations

#### Meetup (draft)

Using Meetup.com's OAuth2 workflow (http://www.meetup.com/meetup_api/auth/#oauth2), easy-megaphone should
run a web server to handle the redirect.

Users will run easy-megaphone, go to the link provided by easy-megaphone to authorize it via OAuth2, where it
redirects back to easy-megaphone running locally and it can then talk to Meetup to post an event.
