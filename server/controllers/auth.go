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

const (
	googleUserInfoEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
)

var conf *oauth2.Config
var randomState string

func init() {
	for _, v := range []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET"} {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("%v must be set in the environment!", v))
		}
	}

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

// HandleLogin sends the user to Google for authentication
func HandleLogin(c *gin.Context) {
	url := conf.AuthCodeURL(randomState)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleCallback parses the access token and exchanges it for the user information
func HandleCallback(c *gin.Context) {
	if state := c.Query("state"); state != randomState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s is not %s", state, randomState))
		return
	}

	// Handle the exchange code to initiate a transport.
	token, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Construct the client.
	client := conf.Client(oauth2.NoContext, token)

	resp, err := client.Get(googleUserInfoEndpoint)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Make sure to close the respond body once this function exits
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Unmarshal our shiny new user from Google
	var u models.User
	if err := json.Unmarshal(data, &u); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Now that we have some information about the user, let's store it to a session
	session := sessions.Default(c)

	session.Set("token", u.Email)
	session.Save()

	// Lastly, redirect the user to the front-end app.
	// TODO::Make this dynamic based on the environment.
	c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/login")
}
