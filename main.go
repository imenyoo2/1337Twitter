package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/vonage/vonage-go-sdk"
	"github.com/vonage/vonage-go-sdk/ncco"
)

func main() {
  scraper := twitterscraper.New()

  tweet := scraper.GetTweets(context.Background(), "1337fil", 1)
  Latest := (<-tweet).Text
  for {
    fmt.Println("checking...")
    tweet = scraper.GetTweets(context.Background(), "1337fil", 1)
    if Latest != (<-tweet).Text {
      fmt.Println("new Tweeet GOGOGOGO")
      call()
      break;
    }
    time.Sleep(5 * time.Second)
  }
}

func call() {
  privateKey, _ := ioutil.ReadFile("/home/nyly/Downloads/private2.key")
  auth, _ := vonage.CreateAuthFromAppPrivateKey("f5b5430c-5322-4f10-b29d-f5d5c2863071", privateKey)
  client := vonage.NewVoiceClient(auth)

  from := vonage.CallFrom{Type: "phone", Number: "212621378024"}
  to := vonage.CallTo{Type: "phone", Number: "212621378024"}

  MyNcco := ncco.Ncco{}
  talk := ncco.TalkAction{Text: "Go library calling to say hello", VoiceName: "Nicole"}
  MyNcco.AddAction(talk)

    // NCCO example
  result, _, _ := client.CreateCall(vonage.CreateCallOpts{From: from, To: to, Ncco: MyNcco})
    // alternate version with answer URL
    //result, _, _ := client.CreateCall(CreateCallOpts{From: from, To: to, AnswerUrl: []string{"https://example.com/answer"}})
  fmt.Println(result.Uuid + " call ID started")
}
