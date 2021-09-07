package gcal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	tokenStash     = "/tmp/gcal-token"
	credentialPath = "./gcal/credentials.json"
)

func GetUpcomingMeeting() {
	client, err := newGcalClient()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	eventService := calendar.NewEventsService(client)

	eventLister := eventService.List("contrbas@gmail.com")

	eventLister.MaxResults(500)
	eventLister.OrderBy("startTime")
	eventLister.SingleEvents(true)
	eventLister.TimeMin(time.Now().Format(time.RFC3339))
	events, err := eventLister.Do()

	if err != nil {
		fmt.Println("error retreiving events", err)
	}
	// fmt.Println("events: ", events)

	for i := 0; i < 500; i++ {
		_, err := fmt.Println(events.Items[i].Summary)
		if err != nil {
			break
		}

	}
	event1 := events.Items[0]
	// event2 := events.Items[1].Summary

	// fmt.Printf("%#v\n", event)

	print(event1.Summary)
	// print(event2)
	if event1.ConferenceData != nil {
		for _, entry := range event1.ConferenceData.EntryPoints {
			if entry.EntryPointType == "video" {
				fmt.Println(entry.Uri)
			}
		}
	}

}

func newGcalClient() (*calendar.Service, error) {
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

	ctx := context.Background()

	return calendar.NewService(context.Background(), option.WithTokenSource(config.TokenSource(ctx, token)))
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
