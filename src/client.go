package logpusher

import (
	"time"
)

const (

	// PostActionURL API endpoint of logpusher
	PostActionURL = "https://api.logpusher.com/api/agent/savelog"
)

// Client all parameter about log or exceptions.
type Client struct {
	authKey     string
	apiKey      string
	logMessage  string
	source      string
	category    string
	logType     string
	logTime     string
	createdDate string
	eventID     string
	email       string
	password    string
}

func New(email, password, apiKey string) Client {
	c := Client{email: email, apiKey: apiKey, password: password}
	return c
}

func (c *Client) Push(message, source, category, logtype, eventid string, logtime time.Timer, date time.Time) (rsp string, err error) {

	return rsp, err
}

func (c *Client) AutoPush(message, source, category, logtype string) (rsp string, err error) {

	return rsp, err
}

func (c *Client) do() error {
	return nil
}
