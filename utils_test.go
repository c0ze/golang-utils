package utils

import "testing"

func TestStripColon(t *testing.T) {
	macAddr := "b8:27:eb:ee:85:2c"
	strippedMacAddr := "B827EBEE852C"

	if StripColon(macAddr) != strippedMacAddr {
		t.Errorf("strip macc address failed " + StripColon(macAddr))
	}
}
