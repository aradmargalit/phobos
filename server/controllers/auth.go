package controllers

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

	// UserID is the constant for the session's user ID
	UserID = "userId"
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
func (e *Env) HandleLogin(c *gin.Context) {
	url := conf.AuthCodeURL(randomState)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleCallback parses the access token and exchanges it for the user information
func (e *Env) HandleCallback(c *gin.Context) {
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

	// Fetch the user information from Google
	resp, err := client.Get(googleUserInfoEndpoint)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Make sure to close the response body once this function exits
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

	// Check if the user exists in the database
	_, err = e.DB.GetUserByEmail(u.Email)

	// If we got an err, there's no user in the database
	if err != nil {
		err = e.DB.InsertUser(u)
		if err != nil {
			panic(err)
		}
	}

	// Now that we have some information about the user, let's store it to a session
	session := sessions.Default(c)
	session.Set(UserID, u.Email)
	session.Save()

	// Lastly, redirect the user to the front-end app.
	// TODO::Make this dynamic based on the environment.
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/login")
}

// Logout will clear the current users cookie
func (e *Env) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserID)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(UserID)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out the user"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/logout")
}
