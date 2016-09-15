package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	server := httptest.NewServer(router()) //starts a real server on a free port
	defer server.Close()                   //notice we use defer here to ensure our server is closed
	req, err := http.NewRequest("POST", server.URL+"/api/echo", strings.NewReader(`{"message":"test"}`))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	echo, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	message := &Message{}
	if err := json.Unmarshal(echo, message); err != nil {
		log.Fatal(err)
	}
	if "martin" != message.Message {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}

func absInt64(x int64) int64 {
	if x < 0 {
		x = -x
	}
	return x
}

func TestTime(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/time", nil)
	if err != nil {
		log.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Time) // the handler we're testing

	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Fail()
		log.Fatal("bad status code returned")
	}

	timeRes, err := ioutil.ReadAll(responseRecorder.Body)
	if err != nil {
		log.Fatal(err)
	}

	messageTime := &MessageTime{}
	if err := json.Unmarshal(timeRes, messageTime); err != nil {
		log.Fatal(err)
	}
	now := time.Now().Unix()
	diffFromNow := absInt64(now - messageTime.Time)
	if diffFromNow > 10 {
		t.Fail()
		log.Println("expected the message to be witint 10 seconds of now")
	}
}
