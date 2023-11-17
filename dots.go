package dots

import (
	"math"
	"strings"
)

var charMap = [4][2]int{
	{1, 8},
	{2, 16},
	{4, 32},
	{64, 128},
}

type Canvas struct {
	Width  int
	Height int
	chars  []int
}

func NewCanvas(width, height int) *Canvas {
	w := int(math.Round(float64(width) / 2.0))
	h := int(math.Round(float64(height) / 4.0))

	chars := make([]int, w*h)
	for i := range chars {
		chars[i] = 0x2800
	}

	return &Canvas{
		Width:  w,
		Height: h,
		chars:  chars,
	}
}

func blockCoord(x, y int) (int, int) {
	return x / 2, y / 4
}

func mapChar(x, y int) int {
	return charMap[y%4][x%2]
}

func (c Canvas) convert(x, y int) int {
	return x + y*c.Width
}

func (c *Canvas) setPixel(x, y int) {
	bX, bY := blockCoord(x, y)
	idx := c.convert(bX, bY)

	c.chars[idx] += mapChar(x, y)
}

func (c *Canvas) Clear() {
	c.chars = make([]int, c.Width*c.Height)
}

func (c *Canvas) String() string {
	var b strings.Builder

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			b.WriteString(string(rune(c.chars[c.convert(x, y)])))
		}
		b.WriteString("\n")
	}

	return b.String()
}
