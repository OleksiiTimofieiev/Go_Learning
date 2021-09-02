package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const (
	tokenStash     = "/tmp/gcal-token"
	credentialPath = "./credentials.json"
)

func GetUpcomingMeeting() {
	_, err := newGcalClient()
	fmt.Println("err: ", err)
}

func newGcalClient() (*http.Client, error) {
	credentialsB, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(credentialsB, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	token, err := getToken(config)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)

	return client, nil
}

func getToken(config *oauth2.Config) (*oauth2.Token, error) {
	stashedToken, err := getStashedToken()
	if err != nil {
		fmt.Println("Error getting stashed token: %s", err)
		token, err := getNewToken(config)
		if err != nil {
			return token, err
		}
		stashToken(token)
		return token, nil

	}
	return stashedToken, nil
}

func getStashedToken() (*oauth2.Token, error) {
	tokenB, err := ioutil.ReadFile(tokenStash)
	if err != nil {
		return &oauth2.Token{}, err
	}
	if len(tokenB) == 0 {
		fmt.Errorf("token file empty")
		return &oauth2.Token{}, err
	}
	var token oauth2.Token
	err = json.Unmarshal(tokenB, &token)

	return &token, nil
}

func stashToken(token *oauth2.Token) error {
	tokenB, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tokenStash, []byte(tokenB), 0644)
}

func getNewToken(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Errorf("Unable to read authorization code: %v", err)
		return &oauth2.Token{}, err
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		fmt.Errorf("Unable to exchange oath code for token from web: %v", err)
		return &oauth2.Token{}, err
	}
	return token, err

}
