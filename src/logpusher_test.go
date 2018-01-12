package logpusher

import "testing"

func TestPush(t *testing.T) {

	//client := logpusher.New("","","")
	client := New("mail@mail.com", "pass", "getapikey")

	if client.apiKey == "" {
		t.Log("done")
	}
}
