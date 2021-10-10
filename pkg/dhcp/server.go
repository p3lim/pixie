package dhcp

import (
	"net"
	"strconv"

	"github.com/p3lim/pixie/pkg/log"
)

type Server struct {
	IP        net.IP
	Port      int
	Broadcast bool
}

func NewServer(addr string) *Server {
	ipStr, portStr, _ := net.SplitHostPort(addr) // already validated
	port, _ := strconv.Atoi(portStr)             // already validated
	ip := net.ParseIP(ipStr)

	if ip == nil {
		// default to broadcast
		ip = net.IPv4zero
	}

	server := &Server{
		IP:   ip,
		Port: port,
	}

	if server.IP.Equal(net.IPv4zero) {
		server.Broadcast = true
	}

	return server
}

func (s *Server) Serve() error {
	conn, err := net.ListenPacket("udp4", net.JoinHostPort(s.IP.String(), strconv.Itoa(s.Port)))
	if err != nil {
		return err
	}
	defer conn.Close()

	return s.loopServe(conn)
}

// DHCPMessageSize defines the maximum DHCP message size supported by this software.
// RFC2132 states that the _minimum_ legal value is 576 octets, but doesn't specify max size, so
// just to be safe we double it.
const DHCPMessageSize int = 1152

func (s *Server) loopServe(conn net.PacketConn) error {
	// we're listening for any messages on the socket
	buffer := make([]byte, DHCPMessageSize)
	for {
		// read from the connection, dumping it in a buffer
		bufLength, sourceAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			return err
		}
		log.Debugf("addr: %v", sourceAddr)

		// RFC2132 states 576 is the minimum DHCP message size, but messages are often shorter than
		// that, I've seen 316 in testing
		// TODO: research further which number we should be looking for, compatibility is important
		if bufLength < 316 {
			// this is not a DHCP message
			continue
		}

		msg, err := ParseMessage(buffer[:bufLength])
		if err != nil {
			return err
		}

		// test the Message for a field that is always there, e.g. hlen

		if msg.GetHLEN() > 16 {
			// hlen can't be more than 16 bytes
			continue
		}

		log.Debugf("op: %v", msg.GetOP())
		log.Debugf("htype: %v", msg.GetHTYPE())
		log.Debugf("hlen: %v", msg.GetHLEN())
		log.Debugf("hops: %v", msg.GetHOPS())
		log.Debugf("xid: %v", msg.GetXID())
		log.Debugf("secs: %v", msg.GetSECS())
		log.Debugf("flag (broadcast): %v", msg.GetFLAGS().Broadcast())
		log.Debugf("ciaddr: %v", msg.GetCIADDR())
		log.Debugf("yiaddr: %v", msg.GetYIADDR())
		log.Debugf("siaddr: %v", msg.GetSIADDR())
		log.Debugf("giaddr: %v", msg.GetGIADDR())
		log.Debugf("chaddr: %v", msg.GetCHADDR())
		log.Debugf("sname: %v", msg.GetSNAME())
		log.Debugf("file: %v", msg.GetFILE())
		log.Debugf("cookie: %v", msg.GetMagicCookie())
		log.Debugf("options: %v", msg.GetOPTIONS())
		log.Debug("------------")
	}
}
