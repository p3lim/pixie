package dhcp

type MessageType byte

const (
	// https://datatracker.ietf.org/doc/html/rfc2132#section-9.6
	Discover MessageType = 1
	Offer    MessageType = 2
	Request  MessageType = 3
	Decline  MessageType = 4
	Ack      MessageType = 5
	Nak      MessageType = 6
	Release  MessageType = 7
	Inform   MessageType = 8
)

func (m MessageType) String() string {
	switch m {
	case Discover:
		return "DHCPDISCOVER"
	case Offer:
		return "DHCPOFFER"
	case Request:
		return "DHCPREQUEST"
	case Decline:
		return "DHCPDECLINE"
	case Ack:
		return "DHCPACK"
	case Nak:
		return "DHCPNAK"
	case Release:
		return "DHCPRELEASE"
	case Inform:
		return "DHCPINFORM"
	default:
		return ""
	}
}
