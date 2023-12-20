package addressmanager

import "github.com/Metchain/MetblockD/appmessage"

// reachabilityFrom returns the relative reachability of the provided local
// address to the provided remote address.
func reachabilityFrom(localAddress, remoteAddress *appmessage.NetAddress, acceptUnroutable bool) int {
	const (
		Unreachable = 0
		Default     = iota
		Teredo
		Ipv6Weak
		Ipv4
		Ipv6Strong
		Private
	)

	IsRoutable := func(na *appmessage.NetAddress) bool {
		if acceptUnroutable {
			return !IsLocal(na)
		}

		return IsValid(na) && !(IsRFC1918(na) || IsRFC2544(na) ||
			IsRFC3927(na) || IsRFC4862(na) || IsRFC3849(na) ||
			IsRFC4843(na) || IsRFC5737(na) || IsRFC6598(na) ||
			IsLocal(na) || (IsRFC4193(na)))
	}

	if !IsRoutable(remoteAddress) {
		return Unreachable
	}

	if IsRFC4380(remoteAddress) {
		if !IsRoutable(localAddress) {
			return Default
		}

		if IsRFC4380(localAddress) {
			return Teredo
		}

		if IsIPv4(localAddress) {
			return Ipv4
		}

		return Ipv6Weak
	}

	if IsIPv4(remoteAddress) {
		if IsRoutable(localAddress) && IsIPv4(localAddress) {
			return Ipv4
		}
		return Unreachable
	}

	/* ipv6 */
	var tunnelled bool
	// Is our v6 is tunnelled?
	if IsRFC3964(localAddress) || IsRFC6052(localAddress) || IsRFC6145(localAddress) {
		tunnelled = true
	}

	if !IsRoutable(localAddress) {
		return Default
	}

	if IsRFC4380(localAddress) {
		return Teredo
	}

	if IsIPv4(localAddress) {
		return Ipv4
	}

	if tunnelled {
		// only prioritise ipv6 if we aren't tunnelling it.
		return Ipv6Weak
	}

	return Ipv6Strong
}
