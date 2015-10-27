package services

import (
  "github.com/Termina1/gogitter"
  // "fmt"
)

type gitterService struct {
  token string
  roomId string
  sent map[string]bool
  serviceName string
}

func NewGitter(token, roomId string) gitterService {
  return gitterService{token, roomId, make(map[string]bool), "gitter"}
}

func (g gitterService) ExposePortal() Portal {
  outGitter, _ := gogitter.GetMessageStream(g.token, g.roomId)
  inGitter := gogitter.GetSendMessageStream(g.token, g.roomId)
  in := make(chan PortalMessage)
  out := make(chan PortalMessage)
  go g.listenToMessages(in, inGitter)
  go g.triggerMessages(out, outGitter)
  return Portal{in, out}
}

func (g gitterService) listenToMessages(in chan PortalMessage, inGitter chan string) {
  for m := range in {
    inGitter <- "**[" + m.Author + "]** " + m.Data
    id := <- inGitter
    g.sent[id] = true
  }
}

func (g gitterService) triggerMessages(out chan PortalMessage, outGitter chan gogitter.GitterMessage) {
  for m := range outGitter {
    _, ok := g.sent[m.Id]
    if !ok {
      message := PortalMessage{m.Text, m.FromUser.DisplayName, g.serviceName, PORTAL_MESSAGE}
      out <- message
    }
  }
}
