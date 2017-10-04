package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"
)

type Pixel struct {
	Red   byte
	Green byte
	Blue  byte
}

func (p *Pixel) AsWord() (w uint16) {
	w = uint16(p.Red & 63)
	w <<= 6
	w += uint16(p.Green & 127)
	w <<= 5
	w += uint16(p.Blue & 63)
	return
}

func (p *Pixel) WriteTo(w io.Writer) (e error) {
	return binary.Write(w, binary.LittleEndian, p.AsWord())
}

func main() {
	r := Pixel{
		Red: 0xff,
	}
	g := Pixel{
		Green: 0xff,
	}
	b := Pixel{
		Blue: 0xff,
	}
	pixels := []Pixel{r, g, b}

	for {
		for _, p := range pixels {
			pWord := p.AsWord()
			fmt.Println(pWord)
			screen := make([]uint16, 320*240)
			for i := 0; i < 320*240; i++ {
				screen[i] = pWord
			}

			f, _ := os.OpenFile("/dev/fb1", os.O_WRONLY, 0)
			binary.Write(f, binary.LittleEndian, screen)
			f.Close()
			<-time.After(time.Second)
		}
	}
}