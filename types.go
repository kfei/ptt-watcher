package main

type Config struct {
	Subscriptions []Subscription `json:"subscriptions"`
	SlackToken    string         `json:"slackToken"`
	SlackChannel  string         `json:"slackChannel"`
	SlackUserName string         `json:"slackUserName"`
}

type Filter string

type Subscription struct {
	Name        string   `json:"name"`
	FeedUrl     string   `json:"feedUrl"`
	RefreshTime int64    `json:"refreshTime"`
	Filters     []Filter `json:"filters"`
}

type Notification struct {
	Body string
}
