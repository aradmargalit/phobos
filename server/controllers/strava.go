package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	models "server/models"
	"server/utils"
	"strconv"
	"strings"
	"time"

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
		RedirectURL:  os.Getenv("SERVER_URL") + "/public/strava/callback",
		Scopes: []string{
			// Need activity:read_all to get private activities
			"activity:read_all",
		},
		Endpoint: stravaAuthorizationEndpoint,
	}

	stravaRandomState = utils.RandomToken()
}

const (
	expiryFormat = "2006-01-02 15:04:05"
	baseURL      = "https://www.strava.com/api/v3"
)

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
	uid, _ := strconv.Atoi(stateParts[1])

	// In the event that the user already has a token in the database, we'll want to update it
	upsertToken(toDatabaseTokens(token, uid), e.DB)

	// Now that we have a token for the user, send them back to the UI
	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"/home")
}

// StravaStatisticsHandler fetches the user's statistics from Strava in real-time
// Todo: deprecate this eventually so I don't exhaust my daily rate limit
func (e *Env) StravaStatisticsHandler(c *gin.Context) {
	// Get a valid client
	client := getHTTPClient(c, e.DB)
	resp, err := client.Get(baseURL + "/athlete")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		c.PureJSON(http.StatusOK, bodyString)
		return
	}

	c.AbortWithError(resp.StatusCode, fmt.Errorf("Error fetching Strava: %v", "but who knows what.."))
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

func getHTTPClient(c *gin.Context, db *models.DB) *http.Client {
	// 1. Pull the user out of the context
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
	}

	// 2. Get the current access and refresh tokens from the DB
	dbToken, err := db.GetStravaTokenByUserID(uid.(int))
	if err != nil {
		panic(err)
	}

	// 3. The OAuth2 library kindly handles refreshes for us as needed. Blessed.
	tokenSource := conf.TokenSource(oauth2.NoContext, toOAuthToken(dbToken))
	newToken, err := tokenSource.Token()
	if err != nil {
		panic(err)
	}

	// 4. Check if the new token is different, and if so, persist that sucker
	if newToken.AccessToken != dbToken.AccessToken {
		fmt.Println("Refresh successful! Updating the user's token!")
		db.UpdateStravaToken(toDatabaseTokens(newToken, uid.(int)))
	}

	return oauth2.NewClient(oauth2.NoContext, tokenSource)
}

func toDatabaseTokens(oauthToken *oauth2.Token, uid int) models.StravaToken {
	formattedExpiry := oauthToken.Expiry.UTC().Format(expiryFormat)

	return models.StravaToken{
		UserID:       uid,
		AccessToken:  oauthToken.AccessToken,
		RefreshToken: oauthToken.RefreshToken,
		Expiry:       formattedExpiry,
	}
}

func toOAuthToken(dbToken models.StravaToken) *oauth2.Token {
	expiresAt, _ := time.Parse(expiryFormat, dbToken.Expiry)
	return &oauth2.Token{
		AccessToken:  dbToken.AccessToken,
		RefreshToken: dbToken.RefreshToken,
		Expiry:       expiresAt,
	}
}
