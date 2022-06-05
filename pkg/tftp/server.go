package tftp

import (
	"bytes"
	"io"
	"net"
	"os"
	"strings"

	"github.com/p3lim/pixie/pkg/log"
	"github.com/pin/tftp"
)

type Server struct {
	Addr string

	chain string
}

func NewServer(addr string, http string) *Server {
	_, httpPort, _ := net.SplitHostPort(http) // already validated
	return &Server{
		Addr:  addr,
		chain: strings.Replace(chainTemplate, "HTTP_PORT", httpPort, 1),
	}
}

func (server *Server) Serve() error {
	tftpServer := tftp.NewServer(server.readHandler, server.writeHandler)
	return tftpServer.ListenAndServe(server.Addr)
}

func (server *Server) readHandler(filename string, rf io.ReaderFrom) (err error) {
	remoteAddr := ""
	if raddr, ok := rf.(tftp.OutgoingTransfer); ok {
		r := raddr.RemoteAddr()
		remoteAddr = r.String()
	}

	log.Debugf("request for '%s' by '%s'", filename, remoteAddr)

	var n int64
	switch filename {
	case "chain.ipxe":
		n, err = rf.ReadFrom(strings.NewReader(server.chain))
	case "undionly.kpxe":
		n, err = rf.ReadFrom(bytes.NewReader(undionly))
	case "ipxe.efi":
		n, err = rf.ReadFrom(bytes.NewReader(ipxe64))
	case "ipxe32.efi":
		n, err = rf.ReadFrom(bytes.NewReader(ipxe32))
	default:
		err = os.ErrNotExist
	}

	log.Debugf("%d bytes written to '%s'", n, remoteAddr)

	return err
}

func (Server) writeHandler(filename string, wt io.WriterTo) error {
	remoteAddr := ""
	if raddr, ok := wt.(tftp.IncomingTransfer); ok {
		r := raddr.RemoteAddr()
		remoteAddr = r.String()
	}

	log.Debugf("request to write '%s' denied for '%s'", filename, remoteAddr)
	return os.ErrPermission
}
