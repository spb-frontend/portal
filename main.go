package main

import (
  "github.com/spb-frontend/portal/services"
  "fmt"
  "encoding/json"
  "os"
)

type GitterConf struct {
  Token string
  Room string
}

type SlackConf struct {
  Channel string
  Token string
}

type Configuration struct {
  Gitter GitterConf
  Slack SlackConf
}

func main() {
  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  conf := Configuration{}
  err := decoder.Decode(&conf)
  if err != nil {
    fmt.Println("error:", err)
  }
  gitter := services.NewGitter(conf.Gitter.Token, conf.Gitter.Room)
  slack := services.NewSlack(conf.Slack.Token, conf.Slack.Channel)
  slackp := slack.ExposePortal()
  gitterp := gitter.ExposePortal()

  for {
    select {
    case m := <- gitterp.Output:
      slackp.Input <- m
    case m := <-slackp.Output:
      gitterp.Input <- m
    }
  }
}
