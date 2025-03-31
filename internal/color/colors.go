package color

var (
	Red         = NewColor(255, 0, 0)
	Orange      = NewColor(255, 126, 0)
	Yellow      = NewColor(255, 255, 0)
	YellowGreen = NewColor(126, 255, 0)
	Green       = NewColor(0, 255, 0)
	GreenCyan   = NewColor(0, 255, 126)
	Cyan        = NewColor(0, 255, 255)
	CyanBlue    = NewColor(0, 126, 255)
	Blue        = NewColor(0, 0, 255)
	BlueMagenta = NewColor(126, 0, 255)
	Magenta     = NewColor(255, 0, 255)
	MagentaRed  = NewColor(255, 0, 126)
)

var Rainbow = []Color{
	Red,
	Orange,
	Yellow,
	YellowGreen,
	Green,
	GreenCyan,
	Cyan,
	CyanBlue,
	Blue,
	BlueMagenta,
	Magenta,
	MagentaRed,
}
