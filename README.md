# search-api

The idea behind the project is to search multiple resources sites(Google Search, Twitter Search, and DuckDuckGO Search) in parallel for a single request made to the API, with the constraint of response timeout. 

Here I have used Context package to timeout request.  the Context is canceled automatically when the timeout elapses. At the current point response returned from the API contains only selected field from the each resource site, but full response can be obtained by updating response model for each API. 


# How To Run

to run this project, set following environment variables.

  - GOOGLE_API_KEY
  - TWITTER_CONSUMER_KEY
  - TWITTER_CONSUMER_SECRET
  - PORT

and then from the main folder run **go run main.go** .

and then intiiate request using cUrl command. example: **curl “http://address:port?q=new+york”**
