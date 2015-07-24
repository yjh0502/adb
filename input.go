package adb

import (
	"image"
	"strconv"
)

func Tap(x, y int) error {
	return TapPoint(image.Pt(x, y))
}

func TapPoint(pt image.Point) error {
	args := []string{strconv.Itoa(pt.X), strconv.Itoa(pt.Y)}
	return execWith(tapCommand(), args)
}

func Back() error {
	return execWith(backCommand(), nil)
}

func Swipe(from, to image.Point) error {
	args := []string{
		strconv.Itoa(from.X),
		strconv.Itoa(from.Y),
		strconv.Itoa(to.X),
		strconv.Itoa(to.Y),
	}
	return execWith(swipeCommand(), args)
}
