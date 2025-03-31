package commands

type GetCommand struct {
	BaseCommand
}

func NewGetCommand() GetCommand {
	return GetCommand{
		BaseCommand: NewBaseCommand("getPilot"),
	}
}
