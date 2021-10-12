package net

import (
	"net"
	"testing"
)

// TestIPString2Int test IPString2Int
func TestIPString2Int(t *testing.T) {
	ipUint, err := IPString2Int("192.168.10.1")
	if err != nil {
		t.Errorf("[TestIPString2Int error] %v", err)
	}
	if ipUint != 3232238081 {
		t.Errorf("[TestIPString2Int error] 192.168.10.1 convert error: %d", ipUint)
	}

	_, err = IPString2Int("392.168.10.1")
	if err == nil {
		t.Error("[TestIPString2Int error] convertion shoud be error")
	}
}

// TestNetIP2Int test NetIP2Int
func TestNetIP2Int(t *testing.T) {
	ip := net.ParseIP("192.168.10.1")

	ipUint := NetIP2Int(ip)
	if ipUint != 3232238081 {
		t.Errorf("[TestIPString2Int error] 192.168.10.1 convert error： %d", ipUint)
	}
}

// TestIPRangeIsOverlap test IPRangeIsOverlap
func TestIPRangeIsOverlap(t *testing.T) {
	ip1 := net.ParseIP("192.169.10.11").To4()
	ip2 := net.ParseIP("192.169.10.20").To4()
	ip3 := net.ParseIP("192.169.10.21").To4()
	ip4 := net.ParseIP("192.169.10.30").To4()
	ip5 := net.ParseIP("192.169.10.15").To4()

	if IPRangeIsOverlap(ip1, ip2, ip3, ip4) {
		t.Error("[TestMinUint error] There is no overlap between [192.169.10.11, 192.169.10.20] and [192.169.10.21, 192.169.10.30]")
	}

	if IPRangeIsOverlap(ip3, ip4, ip1, ip2) {
		t.Error("[TestMinUint error] There is no overlap between [192.169.10.21, 192.169.10.30] and [192.169.10.11, 192.169.10.20]")
	}

	if !IPRangeIsOverlap(ip1, ip2, ip5, ip4) {
		t.Error("[TestMinUint error] There is overlap between [192.169.10.11, 192.169.10.20] and [192.169.10.15, 192.169.10.30]")
	}
}
