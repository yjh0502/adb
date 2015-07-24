package adb

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func buildCommand(baseArg []string, extraArg []string) *exec.Cmd {
	prog := baseArg[0]
	args := make([]string, 0)
	args = append(args, baseArg[1:]...)
	args = append(args, extraArg...)

	return exec.Command(prog, args...)
}

func execWith(baseArg []string, extraArg []string) error {
	cmd := buildCommand(baseArg, extraArg)
	return cmd.Run()
}

func Pkill(packageName string) error {
	return execWith(pkillCommand(), []string{packageName})
}

func Pgrep(packageName string) (int, error) {
	cmd := buildCommand(psCommand(), nil)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, packageName) {
			continue
		}
		splits := strings.Split(line, " ")
		// skip whitespaces
		idx := 1
		for idx < len(splits) && splits[idx] == "" {
			idx += 1
		}
		if idx == len(splits) {
			return 0, nil
		}
		pidStr := splits[idx]
		return strconv.Atoi(pidStr)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return 0, nil
}

func AmStart(packageName string, activityName string) error {
	arg := fmt.Sprintf("%s/%s", packageName, activityName)
	return execWith(amStartCommand(), []string{arg})
}
