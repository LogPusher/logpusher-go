package logpusher

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (

	// PostActionURL API endpoint of logpusher
	PostActionURL = "https://api.logpusher.com/api/agent/savelog"
)

// PushResult model
type PushResult struct {
	Message string `json:"message"`
}

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

// New Get API key from logpusher.com
func New(email, password, apiKey string) Client {

	c := Client{email: email, apiKey: apiKey, password: password}
	return c
}

// Push save log over logpusher
func (c *Client) Push(message, source, category, logtype, eventid string, logtime time.Time, createdate time.Time) (result PushResult, err error) {

	c.logMessage = message
	c.source = source
	c.category = category
	c.logType = logtype
	c.logTime = logtime.Format("hh:MM")
	c.createdDate = createdate.Format("yyyy-mm-ddThh:MM:ssZ:Z")
	c.eventID = eventid

	return c.send()
}

// AutoPush quick push. will Date,Time and EventID fields auto generate by function
func (c *Client) AutoPush(message, source, category, logtype string) (result PushResult, err error) {

	c.logMessage = message
	c.source = source
	c.category = category
	c.logType = logtype
	c.logTime = time.Now().Format("15:04")
	c.createdDate = time.Now().Format("2006-01-02T15:04:05Z07:00")
	c.eventID = fmt.Sprintf("%d", time.Now().Unix())

	return c.send()
}

func (c *Client) send() (result PushResult, err error) {

	body := c.reqValues()
	rspText, err := c.do(PostActionURL, body)

	if err != nil {
		return result, err
	}

	result, err = c.unmarshall(rspText)

	if err != nil {
		return result, err
	}

	return result, err
}

func (c *Client) do(url string, body *bytes.Buffer) (string, error) {

	rsp, err := http.Post(url, "application/json", body)

	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *Client) unmarshall(rspText string) (result PushResult, err error) {

	err = json.Unmarshal([]byte(rspText), &result)

	return result, err
}

func (c *Client) reqValues() *bytes.Buffer {

	currentTime := time.Now().Format("01.02.2006 15:04:05")

	values := map[string]string{
		"AuthKey":     c.generateAuthKey(currentTime),
		"ApiKey":      c.apiKey,
		"LogMessage":  c.logMessage,
		"Source":      c.source,
		"Category":    c.category,
		"LogType":     c.logType,
		"LogTime":     c.logTime,
		"CreatedDate": c.createdDate,
		"EventId":     c.eventID,
	}

	jsonValue, _ := json.Marshal(values)

	return bytes.NewBuffer(jsonValue)
}

// generateAuthKey currentTime format muste be: 01.02.2006 00:00:00
func (c *Client) generateAuthKey(currentTime string) string {

	md5CheckSum := md5.Sum([]byte(c.password))
	payload := fmt.Sprintf("%s|%x|%s", c.email, md5CheckSum, currentTime)

	return base64.StdEncoding.EncodeToString([]byte(payload))
}
