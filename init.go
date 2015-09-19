package adb

import (
	"flag"
)

func DeviceId() string {
	return deviceId
}

var (
	deviceId   string
	remotePath string
)

func init() {
	flag.StringVar(&deviceId, "device", "0b37491303302ead", "device id")
	flag.StringVar(&remotePath, "remote", "", "remote path")
}
