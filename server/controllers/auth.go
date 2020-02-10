package auth

import (
	models "server/models"
	utils "server/utils"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config
var randomState string

func init() {
	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"email", "profile",
		},
		Endpoint: google.Endpoint,
	}

	randomState = utils.RandomToken()
}

// HandleHome serves up the Index
func HandleHome(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(200, `<html><body><a href="/login">Login with Google</a></body></html>`)
}

// HandleLogin sends the user to Google for authentication
func HandleLogin(c *gin.Context) {
	url := conf.AuthCodeURL(randomState)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleCallback parses the access token and exchanges it for the user information
func HandleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != randomState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s is not %s", state, randomState))
		return
	}

	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Construct the client.
	client := conf.Client(oauth2.NoContext, tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var u models.User
	if err := json.Unmarshal(data, &u); err != nil {
		panic(err)
	}

	session := sessions.Default(c)

	session.Set("token", u.Email)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/login")
}
