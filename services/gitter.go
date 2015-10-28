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
  go func() {
    for {
      select {
      case m := <- in:
        g.sendMessage(inGitter, m)
      case m := <-outGitter:
        g.emitMessage(m, out)
      }
    }
  }()
  return Portal{in, out}
}

func (g gitterService) sendMessage(inGitter chan string, m PortalMessage) {
  inGitter <- "**[" + m.Author + "]** " + m.Data
  id := <- inGitter
  g.sent[id] = true
}

func (g gitterService) emitMessage(m gogitter.GitterMessage, out chan PortalMessage) {
  _, ok := g.sent[m.Id]
  if !ok {
    message := PortalMessage{m.Text, m.FromUser.Username, g.serviceName, PORTAL_MESSAGE}
    out <- message
  }
}
