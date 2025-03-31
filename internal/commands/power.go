package commands

type StateParams struct {
	State bool `json:"state"`
}

type OnOffCommand struct {
	BaseCommand
	Params StateParams `json:"params"`
}

func NewOnCommand() OnOffCommand {
	return OnOffCommand{
		BaseCommand: NewBaseCommand("setState"),
		Params: StateParams{
			State: true,
		},
	}
}

func NewOffCommand() OnOffCommand {
	return OnOffCommand{
		BaseCommand: NewBaseCommand("setState"),
		Params: StateParams{
			State: false,
		},
	}
}
