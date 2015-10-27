package services

import (
  "github.com/nlopes/slack"
  "fmt"
)

type slackService struct {
  token string
  channel string
  serviceName string
  api *slack.Client
  rtm *slack.RTM
  userCache map[string]*slack.User
  sent map[string]bool
}

func NewSlack(token, channel string) slackService {
  api := slack.New(token)
  rtm := api.NewRTM()
  go rtm.ManageConnection()
  cache := make(map[string]*slack.User)
  return slackService{token, channel, "slack", api, rtm, cache, make(map[string]bool)}
}

func (sl slackService) ExposePortal() Portal {
  in := make(chan PortalMessage)
  out := make(chan PortalMessage)
  go func() {
    for {
      select {
      case msg := <-sl.rtm.IncomingEvents:
        if msg.Channel == sl.channel {
          sending := sl.triggerMessages(msg)
          if sending.Kind == PORTAL_MESSAGE {
            out <- sending
          }
        }
      case msg := <-in:
        sl.listenToMessages(msg)
      }
    }
  }()
  return Portal{in, out}
}

func (sl slackService) listenToMessages(msg PortalMessage) {
  if (msg.Kind == PORTAL_MESSAGE) {
    params := slack.PostMessageParameters{}
    params.Username = msg.Author
    _, time, err := sl.api.PostMessage(sl.channel, msg.Data, params)
    if err != nil {
      fmt.Println(err)
    }
    sl.sent[time] = true
  }
}

func (sl slackService) triggerMessages(msg slack.RTMEvent) PortalMessage {
  switch ev := msg.Data.(type) {
  case *slack.MessageEvent:
    _, hasMessage := sl.sent[ev.Timestamp]
    if !hasMessage {
      user, ok := sl.userCache[ev.User]
      if !ok {
        info, err := sl.api.GetUserInfo(ev.User)
        if err != nil {
          fmt.Println(err)
        }
        sl.userCache[ev.User] = info
        user = info
      }
      return PortalMessage{ev.Text, user.Name, sl.serviceName, PORTAL_MESSAGE}
    }
  }
  return PortalMessage{"", "", "", PORTAL_NOOP}
}
