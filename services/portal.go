package services

type PortalMessage struct {
  Data string
  Author string
  ServiceName string
  Kind string
}

type Portal struct {
  Input chan PortalMessage
  Output chan PortalMessage
}

type PortalService interface {
  ExposePortal() Portal
}

// func ComposePortals(fs, sn Portal) Portal {
//   nin := make(chan PortalMessage)
//   nout := make(chan PortalMessage)
//   go bindPortals(nin, fs.Output)
//   go bindPortals(sn.Input, fs.Output)
//   go bindPortals(fs.Input, sn.Output)
//   go bindPortals(nin, sn.Output)
//   go bindPortals(sn.Input, nout)
//   go bindPortals(fs.Input, nout)
//   return Portal{nin, nout}
// }

func bindPortals(in, out chan PortalMessage) {
    for m := range out {
      in <- m
    }
}

const PORTAL_MESSAGE string = "portal_message"
const PORTAL_NOOP string = "noop"
