// +build android

package adb

import (
	"strconv"
)

func decompressCapture(data []byte) ([]byte, error) {
	return data, nil
}

func captureCommand() []string { return []string{"/system/bin/screencap"} }
func tapCommand() []string     { return []string{"sh", "/system/bin/input", "tap"} }
func swipeCommand() []string   { return []string{"sh", "/system/bin/input", "swipe"} }
func psCommand() []string      { return []string{"/system/bin/ps"} }
func pkillCommand() []string   { return []string{"sh", "/system/bin/am", "force-stop"} }
func amStartCommand() []string { return []string{"sh", "/system/bin/am", "start"} }

func keyCommand(key int) []string {
	return []string{"sh", "/system/bin/input", "keyevent", strconv.Itoa(key)}
}
func dumpSysCommand(arg string) []string {
	return []string{"dumpsys", arg}
}
