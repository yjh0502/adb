package adb

import (
	"fmt"
	"testing"
)

func TestCapture(t *testing.T) {
	deviceId = "0b37491303302ead"
	remotePath = "jihyun@white"
	tmpDir = "/tmp"

	cmd := captureCommand()
	fmt.Printf("%v", cmd)
}
