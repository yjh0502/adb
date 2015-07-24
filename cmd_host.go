// +build !android

package adb

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strconv"
)

func decompressCapture(data []byte) ([]byte, error) {
	reader := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, gzipReader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// adb shell screencap currupts image, so save to to local file and copy
func captureCommand() []string {
	cmdStr := "fb-adb shell 'screencap | gzip -1 -c'"
	if deviceId != "" {
		cmdStr = fmt.Sprintf("fb-adb -s %s shell \"screencap | gzip -1 -c\"", deviceId)
	}
	var cmd []string
	if remotePath == "" {
		cmd = []string{"bash", "-c", cmdStr}
	} else {
		cmd = []string{"ssh", remotePath, cmdStr}
	}

	return cmd
}

func adbShell() []string {
	var cmd []string
	if deviceId != "" {
		cmd = []string{"adb", "-s", deviceId, "shell"}
	} else {
		cmd = []string{"adb", "shell"}
	}

	if remotePath != "" {
		cmd = append([]string{"ssh", remotePath}, cmd...)
	}
	return cmd
}

func tapCommand() []string {
	return append(adbShell(), "input", "tap")
}
func swipeCommand() []string {
	return append(adbShell(), "input", "swipe")
}
func psCommand() []string {
	return append(adbShell(), "ps")
}
func pkillCommand() []string {
	return append(adbShell(), "am", "force-stop")
}
func amStartCommand() []string {
	return append(adbShell(), "am", "start")
}

func keyCommand(key int) []string {
	return append(adbShell(), "input", "keyevent", strconv.Itoa(key))
}
func dumpSysCommand(arg string) []string {
	return append(adbShell(), "dumpsys", arg)
}
