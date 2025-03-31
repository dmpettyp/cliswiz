package commands

import "github.com/dmpettyp/cliswiz/internal/color"

type ColorParams struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
}

type ColorCommand struct {
	BaseCommand
	Params ColorParams `json:"params"`
}

func NewColorCommand(c color.Color) ColorCommand {
	return ColorCommand{
		BaseCommand: NewBaseCommand("setPilot"),
		Params: ColorParams{
			Red:   c.Red,
			Green: c.Green,
			Blue:  c.Blue,
		},
	}
}
