package tftp

import (
	"bytes"
	"io"
	"net"
	"os"
	"strings"

	"github.com/pin/tftp"
)

type Server struct {
	Addr string

	chain string
}

func NewServer(addr string, http string) *Server {
	_, port, _ := net.SplitHostPort(addr) // already validated
	return &Server{
		Addr:  addr,
		chain: strings.Replace(chainTemplate, "HTTP_PORT", port, 1),
	}
}

func (server *Server) Serve() error {
	tftpServer := tftp.NewServer(server.readHandler, server.writeHandler)
	return tftpServer.ListenAndServe(server.Addr)
}

func (server *Server) readHandler(filename string, rf io.ReaderFrom) (err error) {
	switch filename {
	case "chain.ipxe":
		_, err = rf.ReadFrom(strings.NewReader(server.chain))
	case "undionly.kpxe":
		_, err = rf.ReadFrom(bytes.NewReader(undionly))
	case "ipxe.efi":
		_, err = rf.ReadFrom(bytes.NewReader(ipxe64))
	case "ipxe32.efi":
		_, err = rf.ReadFrom(bytes.NewReader(ipxe32))
	default:
		err = os.ErrNotExist
	}

	return err
}

func (Server) writeHandler(filename string, wt io.WriterTo) error {
	return os.ErrPermission
}
