package controllers

import (
	"fmt"
	"net/http"
	"os"
	models "server/models"
	"server/utils"
	"strconv"
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
	// There must be a better way, but I need to know who the user was when Strava hits my callback endpoint
	// To do this, I'm going to send the user ID with the state, which gets returned to me in the callback
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
	}

	// In order to attach the UID, I need to convert it to a string
	url := stravaConf.AuthCodeURL(stravaRandomState + "|userID:" + strconv.Itoa(uid.(int)))

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// StravaCallbackHandler handles the Strava OAuth2.0 reponse containing the user token
func (e *Env) StravaCallbackHandler(c *gin.Context) {
	state := c.Query("state")
	stateParts := strings.Split(state, "|userID:")
	// First, check to make sure they returned the same random state we sent earlier
	if stateParts[0] != stravaRandomState {
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

	// Now, we want to persist these tokens to the database so that our poor, sweet user doesn't need to reauthenticate
	// 1. We need to figure out which of our users authenticated against Strava
	uid := stateParts[1]

	// convert uid to an int and convert expiry to a MySQL datetime string
	userID, _ := strconv.Atoi(uid)
	formattedExpiry := token.Expiry.UTC().Format("2006-01-02 15:05:05.000")

	stravaToken := models.StravaToken{
		UserID:       userID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       formattedExpiry,
	}

	// In the event that the user already has a token in the database, we'll want to update it
	upsertToken(stravaToken, e.DB)

	c.JSON(200, token)
}

func upsertToken(stravaToken models.StravaToken, db *models.DB) {
	// 1. Check if a token already exists
	_, err := db.GetStravaTokenByUserID(stravaToken.UserID)
	if err != nil {
		// This means there isn't one in the database, so insert one
		_, err = db.InsertStravaToken(stravaToken)
		if err != nil {
			panic(err)
		}
		return
	}

	// 2. If we got here, there's an existing token
	// Refresh it with the retrieved token
	_, err = db.UpdateStravaToken(stravaToken)
	if err != nil {
		panic(err)
	}
	return
}
