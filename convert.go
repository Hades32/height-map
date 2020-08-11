package heightmap

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
)

const sampleSize = 1201

func Convert(r io.Reader, w io.Writer) error {
	img := image.NewGray16(image.Rect(0, 0, sampleSize, sampleSize))
	scanner := bufio.NewScanner(r)
	scanner.Split(HeightData)

	i := 0
	for scanner.Scan() {
		x := i % sampleSize
		y := i / sampleSize
		b := scanner.Bytes()
		height := NormalizedHeight(b)*32
		if height != 0 {
			height += 0xFFFF/10
		}
		img.SetGray16(x, y, color.Gray16{Y: height})
		i++
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return png.Encode(w, img)
}

func NormalizedHeight(b []byte) uint16 {
	return uint16(int(BigEndianSignedInt16(b)) + 32768)
}

func BigEndianSignedInt16(b []byte) int16 {
	return int16(b[0])<<8 | int16(b[1])
}

func HeightData(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if atEOF && len(data) < 2 {
		return 0, nil, errors.New("stream contains cut-off data")
	}
	if len(data) < 2 {
		// Request more data.
		return 0, nil, nil
	}
	return 2, data[0:2], nil
}
