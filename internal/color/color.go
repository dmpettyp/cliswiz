package color

import (
	"math"
)

type Color struct {
	Red, Green, Blue int
}

func NewColor(r, g, b int) Color {
	return Color{
		Red:   clamp(r, 0, 255),
		Green: clamp(g, 0, 255),
		Blue:  clamp(b, 0, 255),
	}
}

func clamp(i, min, max int) int {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

func (c Color) Tint(level int) Color {
	level = clamp(level, 0, 10)

	factor := float64(level) / 10.0

	tint := func(c int) int {
		return c + int(math.Round(float64(255-c)*factor))
	}

	return Color{
		Red:   tint(c.Red),
		Green: tint(c.Green),
		Blue:  tint(c.Blue),
	}
}

func (c Color) Shade(level int) Color {
	level = clamp(level, 0, 10)

	factor := float64(level-10) / 10.0

	shade := func(c int) int {
		return c + int(math.Round(float64(c)*factor))
	}

	return Color{
		Red:   shade(c.Red),
		Green: shade(c.Green),
		Blue:  shade(c.Blue),
	}
}
