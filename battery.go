package adb

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func BatteryLevel() (int, error) {
	cmd := buildCommand(dumpSysCommand("battery"), nil)
	out, err := cmd.Output()

	if err != nil {
		return 0, nil
	}
	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.Split(strings.Trim(line, " "), ": ")
		if len(splits) != 2 {
			continue
		}
		if splits[0] == "level" {
			return strconv.Atoi(splits[1])
		}
	}

	return 0, fmt.Errorf("invalid format for dumpsys battery")
}
