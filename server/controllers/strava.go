package controllers

import (
	"fmt"
	"net/http"
	"os"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var stravaConf *oauth2.Config
var stravaRandomState string

func init() {
	// From https://developers.strava.com/docs/authentication/
	stravaAuthorizationEndpoint := oauth2.Endpoint{
		AuthURL:   "https://www.strava.com/api/v3/oauth/authorize",
		TokenURL:  "https://www.strava.com/api/v3/oauth/token",
		AuthStyle: oauth2.AuthStyleInParams,
	}

	stravaConf = &oauth2.Config{
		ClientID:     os.Getenv("STRAVA_CLIENT_ID"),
		ClientSecret: os.Getenv("STRAVA_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SERVER_URL") + "/strava/callback",
		Scopes: []string{
			// Need activity:read_all to get private activities
			"activity:read_all",
		},
		Endpoint: stravaAuthorizationEndpoint,
	}

	stravaRandomState = utils.RandomToken()
}

// StravaLoginHandler sends the user to Google for authentication
func (e *Env) StravaLoginHandler(c *gin.Context) {
	url := stravaConf.AuthCodeURL(stravaRandomState)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// StravaCallbackHandler handles the Strava OAuth2.0 reponse containing the user token
func (e *Env) StravaCallbackHandler(c *gin.Context) {
	// First, check to make sure they returned the same random state we sent earlier
	if state := c.Query("state"); state != stravaRandomState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s is not %s", state, stravaRandomState))
		return
	}

	// Next, confirm they've accepted access to the activity scope
	if scope := c.Query("scope"); !strings.Contains(scope, "activity:read_all") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("you must accept! You must"))
		return
	}

	// Handle the exchange code to initiate a transport.
	token, err := stravaConf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		panic(err)
	}

	c.JSON(200, token)
}
