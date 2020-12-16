// +build windows,!linux

package gateway

import (
	"fmt"
	"testing"
)

func TestParseWindowsOutput(t *testing.T) {
	output := []byte(`
===========================================================================
Interface List
  7...08 00 27 54 9c b1 ......Intel(R) PRO/1000 MT Desktop Adapter
  1...........................Software Loopback Interface 1
===========================================================================

IPv4 Route Table
===========================================================================
Active Routes:
Network Destination        Netmask          Gateway       Interface  Metric
          0.0.0.0          0.0.0.0         10.0.2.2        10.0.2.15     25
         10.0.2.0    255.255.255.0         On-link         10.0.2.15    281
        10.0.2.15  255.255.255.255         On-link         10.0.2.15    281
       10.0.2.255  255.255.255.255         On-link         10.0.2.15    281
        127.0.0.0        255.0.0.0         On-link         127.0.0.1    331
        127.0.0.1  255.255.255.255         On-link         127.0.0.1    331
  127.255.255.255  255.255.255.255         On-link         127.0.0.1    331
      192.168.3.0    255.255.255.0    192.168.2.254        10.0.2.15     26
        224.0.0.0        240.0.0.0         On-link         127.0.0.1    331
        224.0.0.0        240.0.0.0         On-link         10.0.2.15    281
  255.255.255.255  255.255.255.255         On-link         127.0.0.1    331
  255.255.255.255  255.255.255.255         On-link         10.0.2.15    281
===========================================================================
Persistent Routes:
  None

IPv6 Route Table
===========================================================================
Active Routes:
 If Metric Network Destination      Gateway
  1    331 ::1/128                  On-link
  7    281 fe80::/64                On-link
  7    281 fe80::51bf:99d8:e4a1:7534/128
                                    On-link
  1    331 ff00::/8                 On-link
  7    281 ff00::/8                 On-link
===========================================================================
Persistent Routes:
  None
`)

	gateways := parseWindowsOutput(output, "10.0.2.15")

	fmt.Println(gateways)
}
