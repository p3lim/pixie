package dhcp

import "net"

// https://datatracker.ietf.org/doc/html/rfc2131#section-2
/*
   0                   1                   2                   3
   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |     op (1)    |   htype (1)   |   hlen (1)    |   hops (1)    |
   +---------------+---------------+---------------+---------------+
   |                            xid (4)                            |
   +-------------------------------+-------------------------------+
   |           secs (2)            |           flags (2)           |
   +-------------------------------+-------------------------------+
   |                          ciaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          yiaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          siaddr  (4)                          |
   +---------------------------------------------------------------+
   |                          giaddr  (4)                          |
   +---------------------------------------------------------------+
   |                                                               |
   |                          chaddr  (16)                         |
   |                                                               |
   |                                                               |
   +---------------------------------------------------------------+
   |                                                               |
   |                          sname   (64)                         |
   +---------------------------------------------------------------+
   |                                                               |
   |                          file    (128)                        |
   +---------------------------------------------------------------+
   |                                                               |
   |                          options (variable)                   |
   +---------------------------------------------------------------+
*/

type Message []byte

// GetOP returns the "op code" from the DHCP Message.
// The value can be 1 (for BOOTREQUEST) or 2 (for BOOTREPLY).
func (m Message) GetOP() byte {
	return m[0]
}

// GetHTYPE returns the hardware address type from the DHCP Message.
// The most common value is 1, for ethernet.
func (m Message) GetHTYPE() byte {
	return m[1]
}

// GetHLEN returns the length of the hardware address from the DHCP Message.
// This value is typically 16, for ethernet MAC.
func (m Message) GetHLEN() byte {
	return m[2]
}

// GetHOPS returns the "ops" field from the DHCP Message.
// This is set by the client, typically used by relay agents.
func (m Message) GetHOPS() byte {
	return m[3]
}

// GetXID returns the transaction ID from the DHCP Message.
// This is a random number chosen by the client, used by both client and server to associate
// messages and responses between them.
func (m Message) GetXID() []byte {
	return m[4:8]
}

// GetSECS returns the seconds elapsed since the client began the address/renewal process from the
// DHCP Message.
func (m Message) GetSECS() []byte {
	return m[8:10]
}

// GetFLAGS returns additional flags from the DHCP Message.
// See Flags for more info.
func (m Message) GetFLAGS() Flags {
	return Flags(m[10:12])
}

// GetCIADDR returns the client's IP address from the DHCP Message.
// This is only filled if the client is in BOUND, RENEW, or REBINDING state and can respond to ARP
// requests.
func (m Message) GetCIADDR() net.IP {
	return net.IP(m[12:16])
}

// GetYIADDR returns the server's IP address from the DHCP Message.
func (m Message) GetYIADDR() net.IP {
	return net.IP(m[16:20])
}

// GetSIADDR returns the next-server's IP address from the DHCP Message.
func (m Message) GetSIADDR() net.IP {
	return net.IP(m[20:24])
}

// GetGIADDR returns the relay agent's IP address from the DHCP Message.
func (m Message) GetGIADDR() net.IP {
	return net.IP(m[24:28])
}

// GetCHADDR returns the client's hardware address from the DHCP Message.
// This is typically the ethernet MAC.
func (m Message) GetCHADDR() net.HardwareAddr {
	return net.HardwareAddr(m[28 : 28+m.GetHLEN()])
}

// GetSNAME returns the server's hostname from the DHCP Message.
func (m Message) GetSNAME() string {
	return string(m[44:108])
}

// GetFILE returns the boot file name from the DHCP Message.
func (m Message) GetFILE() string {
	return string(m[108:236])
}

// GetOPTIONS returns the optional parameters from the DHCP Message.
// See Options for more info.
func (m Message) GetOPTIONS() Options {
	return m.parseOptions(m[236:])
}
