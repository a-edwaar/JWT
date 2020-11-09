# JWT in Go

This is a template repo to implement JWT authentication using Go and React

## Logic
To further understand the flow of authentication I'm using check out this [doc](https://hasura.io/blog/best-practices-of-using-jwt-with-graphql/).

## Server

The Go server implements JWT by sending an access token in a login response and a refresh token using an HTTPOnly cookie. It then expects the access token to be provided in the authorisation header in subsequent requests. All the validation of the tokens is done using middleware to avoid boilerplate code in the handlers.

## Client

The React client uses `axiom` for implementing interceptions of any requests made to the server. If there is an access token (stored in memory), expiry time is checked and depending on if its valid or not, the refresh api is automatically called allowing the user to stay signed in despite closing the browser. At that point if a refresh cookie is not present, the user is redirected back to the login page. 

## Database

Currently data is stored in memory on the server. I may move it to use MongoDB in the future.
