// +build !linux,!windows

package gateway

import (
	"errors"
	"net"
	"runtime"
)

func GetGatewaysByInterface(itf *net.Interface) ([]string, error) {
	return nil, errors.New("Not implemented for OS : " + runtime.GOOS)
}
