# A go textual processing server
Primarily I'm just playing around with concurrency/http in Go.

Starts a web server listening on port 8000 for requests at "/" that retrieves
content from supplied URLs and performs text manipulations on it. The current
set is to unique the content and remove integers.

## Build

`cd server && go build`

## Run

`./server`

## Usage
[http://textfiles.com/internet](http://textfiles.com/internet/) contains many links to text-only content. Select a couple and ping the server with them:

`curl -X POST -H "Content-Type: text/plain" -d "http://textfiles.com/internet/aboutmtv.txt,http://textfiles.com/internet/acronyms.txt"`

Should yield the output:
`[IMAGE]WHTSV.CO?odquestin'baklyFrf,h_NpwgRmv-(J)cxD"j/!:LYKBUQP<>=`

## UTF-8
Testing this UTF-8 test page: http://mackerron.com/text/test.utf8.txt should yield the following:

`“Naïvely, thcférgdon£ç”.`

This is a string composed of the following bytes:
[226, 128, 156, 78, 97, 195, 175, 118, 101, 108, 121, 44, 32, 116, 104, 99, 102, 195, 169, 114, 103, 100, 111, 110, 194, 163, 195, 167, 226, 128, 157, 46]
