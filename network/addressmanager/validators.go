package addressmanager

import (
	"github.com/Metchain/MetblockD/appmessage"
	"net"
)

// ipNet returns a net.IPNet struct given the passed IP address string, number
// of one bits to include at the start of the mask, and the total number of bits
// for the mask.
func ipNet(ip string, ones, bits int) net.IPNet {
	return net.IPNet{IP: net.ParseIP(ip), Mask: net.CIDRMask(ones, bits)}
}

// IsIPv4 returns whether or not the given address is an IPv4 address.
func IsIPv4(na *appmessage.NetAddress) bool {
	return na.IP.To4() != nil
}

// IsLocal returns whether or not the given address is a local address.
func IsLocal(na *appmessage.NetAddress) bool {
	return na.IP.IsLoopback() || zero4Net.Contains(na.IP)
}

// IsRFC1918 returns whether or not the passed address is part of the IPv4
// private network address space as defined by RFC1918 (10.0.0.0/8,
// 172.16.0.0/12, or 192.168.0.0/16).
func IsRFC1918(na *appmessage.NetAddress) bool {
	for _, rfc := range rfc1918Nets {
		if rfc.Contains(na.IP) {
			return true
		}
	}
	return false
}

// IsRFC2544 returns whether or not the passed address is part of the IPv4
// address space as defined by RFC2544 (198.18.0.0/15)
func IsRFC2544(na *appmessage.NetAddress) bool {
	return rfc2544Net.Contains(na.IP)
}

// IsRFC3849 returns whether or not the passed address is part of the IPv6
// documentation range as defined by RFC3849 (2001:DB8::/32).
func IsRFC3849(na *appmessage.NetAddress) bool {
	return rfc3849Net.Contains(na.IP)
}

// IsRFC3927 returns whether or not the passed address is part of the IPv4
// autoconfiguration range as defined by RFC3927 (169.254.0.0/16).
func IsRFC3927(na *appmessage.NetAddress) bool {
	return rfc3927Net.Contains(na.IP)
}

// IsRFC3964 returns whether or not the passed address is part of the IPv6 to
// IPv4 encapsulation range as defined by RFC3964 (2002::/16).
func IsRFC3964(na *appmessage.NetAddress) bool {
	return rfc3964Net.Contains(na.IP)
}

// IsRFC4193 returns whether or not the passed address is part of the IPv6
// unique local range as defined by RFC4193 (FC00::/7).
func IsRFC4193(na *appmessage.NetAddress) bool {
	return rfc4193Net.Contains(na.IP)
}

// IsRFC4380 returns whether or not the passed address is part of the IPv6
// teredo tunneling over UDP range as defined by RFC4380 (2001::/32).
func IsRFC4380(na *appmessage.NetAddress) bool {
	return rfc4380Net.Contains(na.IP)
}

// IsRFC4843 returns whether or not the passed address is part of the IPv6
// ORCHID range as defined by RFC4843 (2001:10::/28).
func IsRFC4843(na *appmessage.NetAddress) bool {
	return rfc4843Net.Contains(na.IP)
}

// IsRFC4862 returns whether or not the passed address is part of the IPv6
// stateless address autoconfiguration range as defined by RFC4862 (FE80::/64).
func IsRFC4862(na *appmessage.NetAddress) bool {
	return rfc4862Net.Contains(na.IP)
}

// IsRFC5737 returns whether or not the passed address is part of the IPv4
// documentation address space as defined by RFC5737 (192.0.2.0/24,
// 198.51.100.0/24, 203.0.113.0/24)
func IsRFC5737(na *appmessage.NetAddress) bool {
	for _, rfc := range rfc5737Net {
		if rfc.Contains(na.IP) {
			return true
		}
	}

	return false
}

// IsRFC6052 returns whether or not the passed address is part of the IPv6
// well-known prefix range as defined by RFC6052 (64:FF9B::/96).
func IsRFC6052(na *appmessage.NetAddress) bool {
	return rfc6052Net.Contains(na.IP)
}

// IsRFC6145 returns whether or not the passed address is part of the IPv6 to
// IPv4 translated address range as defined by RFC6145 (::FFFF:0:0:0/96).
func IsRFC6145(na *appmessage.NetAddress) bool {
	return rfc6145Net.Contains(na.IP)
}

// IsRFC6598 returns whether or not the passed address is part of the IPv4
// shared address space specified by RFC6598 (100.64.0.0/10)
func IsRFC6598(na *appmessage.NetAddress) bool {
	return rfc6598Net.Contains(na.IP)
}

// IsValid returns whether or not the passed address is valid. The address is
// considered invalid under the following circumstances:
// IPv4: It is either a zero or all bits set address.
// IPv6: It is either a zero or RFC3849 documentation address.
func IsValid(na *appmessage.NetAddress) bool {
	// IsUnspecified returns if address is 0, so only all bits set, and
	// RFC3849 need to be explicitly checked.
	return na.IP != nil && !(na.IP.IsUnspecified() ||
		na.IP.Equal(net.IPv4bcast))
}
