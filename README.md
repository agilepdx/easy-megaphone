# easy-megaphone
Take in JSON, talk to publication services

## Outline

1. Take in JSON file (format TBD).
2. Output to various services

Services should include:
* Meetup.com API
* Twitter API
* Calagator (done)
* AgilePDX website via GitHub Pages (in progress)

## Credentials

Store in environment.

Github integration relies on a [personal access token](https://developer.github.com/v3/oauth/).

`cat ~/github-personal-access-token.sh`

```
#!/bin/sh

export EASYMEGAPHONE_GITHUBTOKEN="token here"
```
