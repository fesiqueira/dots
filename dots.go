package dots

import (
	"math"
	"strings"
)

const brailleOffset rune = 0x2800

var charMap = [4][2]rune{
	{1, 8},
	{2, 16},
	{4, 32},
	{64, 128},
}

type Canvas struct {
	Width  int
	Height int
	width  int
	height int
	chars  []rune
}

func NewCanvas(width, height int) *Canvas {
	w := int(math.Ceil(float64(width) / 2))
	h := int(math.Ceil(float64(height) / 4))

	return &Canvas{
		Width:  width,
		Height: height,
		width:  w,
		height: h,
		chars:  make([]rune, w*h),
	}
}

func getChar(x, y int) rune {
	return charMap[y%4][x%2]
}

func (c Canvas) getBlockIdx(x, y int) int {
	return x + y*c.width
}

func (c Canvas) getIdx(x, y int) int {
	return x/2 + y/4*c.width
}

func (c Canvas) hasPixel(x, y int) bool {
	idx := c.getIdx(x, y)
	return c.chars[idx]&getChar(x, y) > 0
}

func (c *Canvas) togglePixel(x, y int) {
	if c.hasPixel(x, y) {
		c.unsetPixel(x, y)
		return
	}
	c.setPixel(x, y)
}

func (c *Canvas) unsetPixel(x, y int) {
	idx := c.getIdx(x, y)
	c.chars[idx] ^= getChar(x, y)
}

func (c *Canvas) setPixel(x, y int) {
	idx := c.getIdx(x, y)
	c.chars[idx] |= getChar(x, y)
}

func (c Canvas) rows() []string {
	rows := make([]string, c.height)

	for y := 0; y < c.height; y++ {
		row := make([]rune, c.width)
		for x := 0; x < c.width; x++ {
			blk := c.getBlockIdx(x, y)
			row[x] = brailleOffset + c.chars[blk]
		}
		rows[y] = string(row)
	}

	return rows
}

func (c *Canvas) Clear() {
	c.chars = make([]rune, c.width*c.height)
}

func (c *Canvas) String() string {
	var b strings.Builder

	for _, row := range c.rows() {
		b.WriteString(row + "\n")
	}

	return b.String()
}
