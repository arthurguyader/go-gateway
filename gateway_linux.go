// +build linux,!windows

package gateway

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func GetGatewaysByInterface(itf *net.Interface) ([]string, error) {
	link, err := netlink.LinkByName(itf.Name)

	if err != nil {
		return nil, fmt.Errorf("netlink.LinkByName failed <- %v", err)
	}

	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)

	if err != nil {
		return nil, fmt.Errorf("netlink.RouteList failed <- %v", err)
	}

	var gateways []string

	for _, route := range routes {
		if route.Gw != nil {
			gateways = append(gateways, route.Gw.String())
		}

	}
	return gateways, nil
}
