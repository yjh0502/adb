package adb

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"
)

func TurnOffScreen() error {
	if err := setScreenState(false); err != nil {
		return fmt.Errorf("TurnOffScreen: %s", err)
	}
	return nil
}

func TurnOnScreen() error {
	if err := setScreenState(true); err != nil {
		return fmt.Errorf("TurnOnScreen: %s", err)
	}
	return nil
}

func setScreenState(on bool) error {
	state, err := getScreenState()
	if err != nil {
		return err
	}
	if state != on {
		return toggleScreen()
	}
	return nil
}

func getScreenState() (bool, error) {
	cmd := buildCommand(dumpSysCommand("power"), nil)
	out, err := cmd.Output()

	if err != nil {
		return true, nil
	}
	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Display Power") {
			continue
		}

		if strings.Contains(line, "state=ON") {
			return true, nil
		} else if strings.Contains(line, "state=OFF") {
			return false, nil
		}
	}

	return getScreenStateInputMethod()
}

func getScreenStateInputMethod() (bool, error) {
	cmd := buildCommand(dumpSysCommand("input_method"), nil)
	out, err := cmd.Output()

	if err != nil {
		return true, nil
	}
	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "mScreenOn=") {
			continue
		}

		if strings.Contains(line, "mScreenOn=true") {
			return true, nil
		} else if strings.Contains(line, "mScreenOn=false") {
			return false, nil
		}
	}

	return true, fmt.Errorf("invalid format for dumpsys input_method")
}

func toggleScreen() error {
	err := execWith(powerCommand(), nil)
	// wait until screen toggles
	time.Sleep(time.Second)
	return err
}
