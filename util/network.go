package util

import (
	"net"
	"strings"
	"syscall"
	"errors"
)

func CreateListener(address string) (net.Listener, error) {
	dsn := strings.Split(address, "://")
	if len(dsn) != 2 {
		return nil, errors.New("invalid socket DSN (tcp://:6001, unix://file.sock)")
	}

	if dsn[0] == "unix" {
		syscall.Unlink(dsn[1])
	}

	return net.Listen(dsn[0], dsn[1])
}