package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"server/internal/models"
	"server/internal/repository"
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

// HandleStravaLogin sends the user to Strava for authentication
func (svc *service) HandleStravaLogin(c *gin.Context) {
	// There must be a better way, but I need to know who the user was when Strava hits my callback endpoint
	// To do this, I'm going to send the user ID with the state, which gets returned to me in the callback
	uid := c.GetInt("user")

	// In order to attach the UID, I need to convert it to a string
	url := stravaConf.AuthCodeURL(stravaRandomState + "|userID:" + strconv.Itoa(uid))

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleStravaCallback handles the Strava OAuth2.0 reponse containing the user token
func (svc *service) HandleStravaCallback(c *gin.Context) {
	state := c.Query("state")
	stateParts := strings.Split(state, "|userID:")
	// First, check to make sure they returned the same random state we sent earlier
	if stateParts[0] != stravaRandomState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s is not %s", state, stravaRandomState))
		return
	}

	// Next, confirm they've accepted access to the activity scope
	// TODO This should return to the UI with some sort of error
	if scope := c.Query("scope"); !strings.Contains(scope, "activity:read_all") {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"/error/strava")
		return
	}

	// Handle the exchange code to initiate a transport.
	token, err := stravaConf.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		panic(err)
	}

	// The returned token tells us the athlete's ID in the "extra" portion of the response
	athleteInfo := token.Extra("athlete")
	athleteMap, ok := athleteInfo.(map[string]interface{})
	if !ok {
		panic("could not convert athlete information to a map")
	}
	stravaID := int(athleteMap["id"].(float64))

	// Now, we want to persist these tokens to the database so that our poor, sweet user doesn't need to reauthenticate
	// We need to figure out which of our users authenticated against Strava
	uid, _ := strconv.Atoi(stateParts[1])

	// In the event that the user already has a token in the database, we'll want to update it
	upsertToken(toDatabaseTokens(token, uid, stravaID), svc.db)

	// Now that we have a token for the user, send them back to the UI
	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"/home")
}

// HandleStravaDeauthorization is responsible for disconnecting the user from Strava updates
func (svc *service) HandleStravaDeauthorization(uid int) error {

	client := getHTTPClient(uid, svc.db)
	fmt.Println("Deauthorizing user id: " + strconv.Itoa(uid) + " from Strava access...")

	resp, err := client.Post("https://www.strava.com/oauth/deauthorize", "application/json", nil)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Now that we've done that, we need to delete that token from the DB
	err = svc.db.DeleteStravaTokenByUserID(uid)
	if err != nil {
		return err
	}
	return nil
}

func upsertToken(stravaToken models.StravaToken, db repository.PhobosDB) {
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
}

func getHTTPClient(uid int, db repository.PhobosDB) *http.Client {
	// 1. Get the current access and refresh tokens from the DB
	dbToken, err := db.GetStravaTokenByUserID(uid)
	if err != nil {
		panic(err)
	}

	// 2. The OAuth2 library kindly handles refreshes for us as needed. Blessed.
	tokenSource := stravaConf.TokenSource(context.Background(), toOAuthToken(dbToken))
	client := oauth2.NewClient(context.Background(), tokenSource)
	newToken, err := tokenSource.Token()
	if err != nil {
		panic(err)
	}

	// 3. Check if the new token is different, and if so, persist that sucker
	if newToken.AccessToken != dbToken.AccessToken {
		fmt.Println("Refresh successful! Updating the user's token!")
		db.UpdateStravaToken(toDatabaseTokens(newToken, uid, dbToken.StravaID))
	}

	return client
}

func toDatabaseTokens(oauthToken *oauth2.Token, uid int, stravaID int) models.StravaToken {
	formattedExpiry := oauthToken.Expiry.UTC().Format(expiryFormat)

	return models.StravaToken{
		UserID:       uid,
		StravaID:     stravaID,
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
