package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/alexflint/go-arg"
	"github.com/p3lim/pixie/pkg/http"
	"github.com/p3lim/pixie/pkg/log"
	"github.com/p3lim/pixie/pkg/tftp"
)

var Version = "dev"

type args struct {
	TFTP      string `arg:"-t,env" default:"0.0.0.0:69" placeholder:"ADDR" help:"tftp server address and port"`
	HTTP      string `arg:"-u,env" default:"0.0.0.0:80" placeholder:"ADDR" help:"http server address and port"`
	Scripts   string `arg:"-d,env,required" placeholder:"DIR" help:"path to iPXE scripts served from HTTP"`
	Extra     string `arg:"-e,env" placeholder:"DIR" help:"path to extra files served from HTTP"`
	Verbosity string `arg:"-v,env:LEVEL" default:"INFO" placeholder:"LEVEL" help:"verbosity level (ERROR,WARNING,INFO,DEBUG)"`
}

func (args) Version() string {
	return fmt.Sprintf("pixie %s", Version)
}

func main() {
	var cfg args
	arg.MustParse(&cfg)

	if err := cfg.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if level, err := log.GetLevel(cfg.Verbosity); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		log.SetLevel(level)
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		tftpServer := tftp.NewServer(cfg.TFTP, cfg.HTTP)
		log.Infof("tftp server listening on %s", cfg.TFTP)
		log.Fatal(tftpServer.Serve())
		wg.Done()
	}()

	go func() {
		httpServer := http.NewServer(cfg.HTTP, cfg.Scripts, cfg.Extra)
		log.Infof("http server listening on %s", cfg.HTTP)
		log.Fatal(httpServer.Serve())
		wg.Done()
	}()

	wg.Wait()
}
