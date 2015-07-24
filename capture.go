package adb

import (
	"bytes"
	"encoding/binary"
	"image"
	"os/exec"
	"sync"
	"time"
)

var (
	lastImage    *image.NRGBA
	captureMutex sync.Mutex
)

func ScreenCapture() (*image.NRGBA, error) {
	captureMutex.Lock()
	defer captureMutex.Unlock()

	args := captureCommand()
	cmd := exec.Command(args[0], args[1:]...)

	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	data, err := decompressCapture(buf.Bytes())
	if err != nil {
		return nil, err
	}

	buf = bytes.NewBuffer(data)
	width := int32(0)
	height := int32(0)
	version := int32(0)

	if err := binary.Read(buf, binary.LittleEndian, &width); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &height); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &version); err != nil {
		return nil, err
	}

	stride := int(width * 4)
	rect := image.Rectangle{image.Pt(0, 0), image.Pt(int(width), int(height))}

	img := &image.NRGBA{Pix: data[12:], Stride: stride, Rect: rect}

	lastImage = img
	return lastImage, nil
}

type ScreenImage struct {
	Image   *image.NRGBA
	Created time.Time
}

func ScreensCapture(out chan<- ScreenImage, done <-chan bool) error {
	onScreen := make(chan *image.NRGBA)
	onError := make(chan error)
	startTime := time.Now()
	captureFunc := func() {
		img, err := ScreenCapture()
		if err != nil {
			onError <- err
		} else {
			onScreen <- img
		}
	}
	go captureFunc()

	for {
		select {
		case <-done:
			return nil
		case err := <-onError:
			return err
		case img := <-onScreen:
			out <- ScreenImage{
				Image:   img,
				Created: startTime,
			}
			startTime = time.Now()
			go captureFunc()
		}
	}
}
