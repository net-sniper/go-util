package net

import (
	"errors"
	"net"

	"github.com/nutscloud/go-util/math"
)

// IPString2Int Converts the IP string to a number
func IPString2Int(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// NetIP2Int Converts the net.IP to a number
func NetIP2Int(ip net.IP) uint {
	ipv4 := ip.To4()
	return uint(ipv4[3]) | uint(ipv4[2])<<8 | uint(ipv4[1])<<16 | uint(ipv4[0])<<24
}

func IPRangeIsOverlap(aStart, aEnd, bStart, bEnd net.IP) bool {
	aStartInt := NetIP2Int(aStart)
	aEndInt := NetIP2Int(aEnd)
	bStartInt := NetIP2Int(bStart)
	bEndInt := NetIP2Int(bEnd)

	return math.IsRangeOverlapUint(aStartInt, aEndInt, bStartInt, bEndInt)
}
