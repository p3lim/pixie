package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func parseHostPort(hostport string) error {
	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return err
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("unable to parse port")
	}
	if portInt < 1 || portInt > 65535 {
		return fmt.Errorf("invalid port %d", portInt)
	}
	if host != "" && net.ParseIP(host) == nil {
		return fmt.Errorf("invalid ip %s", host)
	}

	return nil
}

func parseDir(path string) error {
	if fs, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist at %s", path)
	} else if os.IsPermission(err) {
		return fmt.Errorf("path is not accessible at %s", path)
	} else if !fs.IsDir() {
		return fmt.Errorf("path is not a directory at %s", path)
	}

	return nil
}

func (cfg args) Validate() error {
	if err := parseHostPort(cfg.TFTP); err != nil {
		return err
	}
	if err := parseHostPort(cfg.HTTP); err != nil {
		return err
	}
	if cfg.DHCP != "" {
		if err := parseHostPort(cfg.DHCP); err != nil {
			return err
		}
	}
	if err := parseDir(cfg.Scripts); err != nil {
		return err
	}
	if cfg.Extra != "" {
		if err := parseDir(cfg.Extra); err != nil {
			return err
		}
	}

	return nil
}
