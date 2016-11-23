package main

type Config struct {
	Subscriptions []Subscription `json:"subscriptions"`
	Notifications Notifications  `json:"notifications"`
}

type Filter string

type Notifications struct {
	Slack SlackConfig `json:"slack,omitempty"`
	Line  LineConfig  `json:"line,omitempty"`
}

type SlackConfig struct {
	Token    string `json:"token"`
	Channel  string `json:"channel"`
	UserName string `json:"userName"`
}

type LineConfig struct {
	ChannelSecret      string `json:"channelSecret"`
	ChannelAccessToken string `json:"channelAccessToken"`
	ToUserId           string `json:"toUserId"`
}

type Subscription struct {
	Name          string   `json:"name"`
	FeedUrl       string   `json:"feedUrl"`
	RefreshTime   int64    `json:"refreshTime"`
	Filters       []Filter `json:"filters"`
	NotifyMethods []string `json:"notifyMethods"`
}

type NotificationMessage struct {
	Body string
}
