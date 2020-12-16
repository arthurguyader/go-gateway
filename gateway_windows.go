// +build windows,!linux

package gateway

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func parseWindowsOutput(output []byte, ip string) []string {
	lines := strings.Split(string(output), "\n")

	gateways := make([]string, 0)
	ipv4RouteTable := make([]string, 0)

	found := false
	sep := 0
	for _, line := range lines {
		if strings.Contains(line, "=") {
			sep++
		}

		if sep == 4 {
			break
		}

		if found {
			ipv4RouteTable = append(ipv4RouteTable, line)
		}

		if sep == 3 {
			found = true
		}
	}

	for _, routeTableLine := range ipv4RouteTable[2:] {
		fields := strings.Fields(routeTableLine)

		if fields[3] == ip {
			gtw := net.ParseIP(fields[2])
			if gtw != nil {
				gateways = append(gateways, gtw.String())
			}
		}
	}

	return gateways
}

func GetGatewaysByInterface(itf *net.Interface) ([]string, error) {
	cmdOutput, err := exec.Command("route", "print").Output()
	if err != nil {
		return nil, err
	}

	addrs, err := itf.Addrs()
	if err != nil {
		return nil, fmt.Errorf("itf.Addrs failed <- %v", err)
	}

	var ip net.IP

	for _, addr := range addrs {

		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		// check if ipv4
		if ip == nil || ip.To4() == nil {
			continue
		}

		return parseWindowsOutput(cmdOutput, ip.String()), nil
	}

	return nil, errors.New("no valid ipv4 address")
}
