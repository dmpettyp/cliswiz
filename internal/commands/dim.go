package commands

type DimParams struct {
	Dimming int `json:"dimming"`
}

type DimCommand struct {
	BaseCommand
	Params DimParams `json:"params"`
}

func NewDimCommand(level int) DimCommand {
	if level < 10 {
		level = 10
	}

	if level > 100 {
		level = 100
	}

	return DimCommand{
		BaseCommand: NewBaseCommand("setPilot"),
		Params: DimParams{
			Dimming: level,
		},
	}
}
