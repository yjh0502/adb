package adb

import (
	"image"
	"strconv"
)

// TODO: landscape-left와 landscape-right의 좌표 translation이 다릅니다
func translate(pt image.Point) image.Point {
	w := lastImage.Bounds().Max.X
	return image.Pt(pt.Y, w-pt.X)
}

func Tap(x, y int) error {
	return TapPoint(image.Pt(x, y))
}

func TapPoint(pt image.Point) error {
	trans := translate(pt)

	args := []string{strconv.Itoa(trans.X), strconv.Itoa(trans.Y)}
	return execWith(tapCommand(), args)
}

func Back() error {
	return execWith(backCommand(), nil)
}

func Swipe(from, to image.Point) error {
	transFrom := translate(from)
	transTo := translate(to)

	args := []string{
		strconv.Itoa(transFrom.X),
		strconv.Itoa(transFrom.Y),
		strconv.Itoa(transTo.X),
		strconv.Itoa(transTo.Y),
	}
	return execWith(swipeCommand(), args)
}
