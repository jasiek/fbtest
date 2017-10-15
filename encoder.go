package main

import (
	"encoding/binary"
	"image"
	"image/color"
	"io"
)

type encoder struct {
	Width  uint
	Height uint
	pixels []uint16
}

func newEncoder(width uint, height uint) (e *encoder) {
	return &encoder{
		Width:  width,
		Height: height,
		pixels: make([]uint16, width*height),
	}
}

func mapColor(v uint32, b uint8) uint16 {
	if b > 8 {
		panic("more than 8 bits are not supported")
	}

	v >>= 6

	buckets := uint16(1 << b)
	return uint16(v) / buckets
}

func convertColor(c color.Color) uint16 {
	model := color.RGBAModel
	c = model.Convert(c)

	val := uint16(0)
	red, green, blue, _ := c.RGBA()

	val += mapColor(red, 5) * (1 << 11)
	val += mapColor(green, 6) * (1 << 6)
	val += mapColor(blue, 5)
	return val
}

func (e *encoder) Encode(w io.Writer, m image.Image) error {
	xStart := m.Bounds().Min.X
	yStart := m.Bounds().Min.Y

	// For simplicity, only use the first 320x240 pixels from the image
	for x := uint(0); x < e.Width; x++ {
		for y := uint(0); y < e.Height; y++ {
			imageX := xStart + int(x)
			imageY := yStart + int(y)

			color := m.At(imageX, imageY)
			e.pixels[y*e.Width+x] = convertColor(color)
		}
	}

	return binary.Write(w, binary.LittleEndian, e.pixels)
}

func Encode(w io.Writer, m image.Image) error {
	return newEncoder(320, 240).Encode(w, m)
}
