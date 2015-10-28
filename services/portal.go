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

const PORTAL_MESSAGE string = "portal_message"
const PORTAL_NOOP string = "noop"
