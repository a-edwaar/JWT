# JWT in Go

This is a boostrap repo to implement JWT authentication using Go and React

## Server

The Go server implements an access token and refresh token using cookies.
Validation of the tokens is done using middleware to avoid boilerplate code in the handlers.

## Client

The React client uses `axiom` for implementing interceptions of any requests made to the server. If there is an access token, expiry time is checked and depending on if its valid or not, the refresh api is automatically called allowing the user to stay signed in despite closing the browser.

## Database

Currently data is stored in memory on the server. I may move it to use MongoDB in the future.
