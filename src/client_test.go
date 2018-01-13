package logpusher

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	returnMessage = `{
		"message": "Object reference not set to an instance of an object."
	}`

	client       Client
	testEmail    = "mail@mail.com"
	testPassword = "p@s"
	testAPIkey   = ""
	testAuthKey  = "bWFpbEBtYWlsLmNvbXwzZjQ3MzdkNzE1NmM1OGNkMjhhYTQzZTAzZmIyZDRhY3wwMS4wMi4yMDA2IDAwOjAwOjAw"
)

func TestGenerateAuthKey(t *testing.T) {
	client = New(testEmail, testPassword, testAPIkey)

	currentTime := "01.02.2006 00:00:00"
	authkey := client.generateAuthKey(currentTime)

	if authkey != testAuthKey {
		t.Error("Incorrect authentication hash", "authkey", authkey, "test authkey", testAuthKey)
	}

}

func TestUnmarshallPushResult(t *testing.T) {
	client = New(testEmail, testPassword, testAPIkey)

	obj, err := client.unmarshall(returnMessage)

	if err != nil {
		t.Error(err)
	}

	if obj.Message != "Object reference not set to an instance of an object." {
		t.Error("Invalid unmarshall PushResult modal.")
	}
}

func TestPostRequest(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		postData := string(body)

		t.Log("Post Data:", postData)
		io.WriteString(w, returnMessage)
	}))

	defer ts.Close()

	client = New(testEmail, testPassword, testAPIkey)
	client.apiKey = testAPIkey
	client.authKey = testAuthKey
	client.category = "billing"
	client.eventID = "000001"
	client.logMessage = "this is exception"
	client.logTime = "14:00"
	client.logTime = "2016-10-24T21:26:46.4000402+03:00"
	client.source = "billing app"

	body := client.reqValues()
	result, err := client.do(ts.URL, body)

	if err != nil {
		t.Error(err)
	}

	if result == "" {
		t.Error("Result cannot be empty")
	}
}
