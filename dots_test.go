package dots

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLine(t *testing.T) {
	c := NewCanvas(10, 10)

	c.Line(Point{0, 0}, Point{9, 0})
	c.Line(Point{0, 0}, Point{0, 9})
	c.Line(Point{9, 0}, Point{9, 9})
	c.Line(Point{0, 9}, Point{9, 9})
	c.Line(Point{0, 0}, Point{9, 9})
	c.Line(Point{0, 9}, Point{9, 0})
	c.Line(Point{0, 4}, Point{9, 4})
	c.Line(Point{4, 0}, Point{4, 9})

	fmt.Printf("%s\n", c)
}

func TestTriangle(t *testing.T) {
	c := NewCanvas(150, 100)

	c.Triangle(Point{130, 80}, Point{1, 20}, Point{30, 10})
	c.Triangle(Point{1, 1}, Point{20, 90}, Point{100, 80})

	fmt.Print(c)
}

func TestCanvas(t *testing.T) {
	for _, tc := range []struct {
		name   string
		width  int
		height int
		points []Point
		want   *Canvas
	}{
		{
			name:   "out of bound",
			width:  1,
			height: 1,
			points: []Point{
				{X: -1, Y: -1},
				{X: 1, Y: 1},
				{X: 10, Y: 10},
			},
			want: &Canvas{
				chars: []rune{16},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCanvas(tc.width, tc.height)

			for _, p := range tc.points {
				c.setPixel(p.X, p.Y)
			}

			if !reflect.DeepEqual(tc.want.chars, c.chars) {
				t.Errorf("want %v, but got %v", tc.want.chars, c.chars)
			}
		})
	}
}

func TestString(t *testing.T) {
	c := NewCanvas(10, 10)

	c.setPixel(2, 3)
	c.setPixel(3, 3)
	c.setPixel(1, 4)
	c.setPixel(1, 5)
	c.setPixel(2, 6)
	c.setPixel(3, 7)
	c.setPixel(4, 8)
	c.setPixel(5, 7)
	c.setPixel(6, 6)
	c.setPixel(7, 5)
	c.setPixel(7, 4)
	c.setPixel(6, 3)
	c.setPixel(5, 3)
	c.togglePixel(4, 4)
	fmt.Printf("%s\n", c)
}

func TestPrintBlock(t *testing.T) {
	c := NewCanvas(10, 10)

	c.setPixel(0, 0)
	c.setPixel(1, 0)
	c.setPixel(2, 0)
	c.setPixel(3, 0)
	c.setPixel(4, 0)
	c.setPixel(5, 0)
	c.setPixel(6, 0)
	c.setPixel(7, 0)
	c.setPixel(8, 0)
	c.setPixel(9, 0)

	c.setPixel(0, 1)
	c.setPixel(0, 2)
	c.setPixel(0, 3)
	c.setPixel(0, 4)
	c.setPixel(0, 5)
	c.setPixel(0, 6)
	c.setPixel(0, 7)
	c.setPixel(0, 8)
	c.setPixel(0, 9)

	c.setPixel(1, 9)

	fmt.Printf("%s\n", c)
}

func TestIdempotency(t *testing.T) {
	c := NewCanvas(1, 1)

	for _, tc := range []struct {
		name    string
		pixelFn func(x, y int)
		want    []rune
	}{
		{name: "set pixel", pixelFn: c.setPixel, want: []rune{16}},
		{name: "unset pixel", pixelFn: c.unsetPixel, want: []rune{0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.pixelFn(1, 1)
			tc.pixelFn(1, 1)
			if !reflect.DeepEqual(c.chars, tc.want) {
				t.Errorf("want `%v` but got `%v`", tc.want, c.chars)
			}
		})
	}
}

func TestTogglePixel(t *testing.T) {
	c := NewCanvas(1, 1)

	want := []rune{16}
	c.togglePixel(1, 1)

	if !reflect.DeepEqual(c.chars, want) {
		t.Errorf("want `%v` but got `%v`", want, c.chars)
	}

	want = []rune{0}
	c.togglePixel(1, 1)

	if !reflect.DeepEqual(c.chars, want) {
		t.Errorf("want `%v` but got `%v`", want, c.chars)
	}
}
