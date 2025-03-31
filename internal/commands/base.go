package commands

type BaseCommand struct {
	ID     string `json:"id"`
	Method string `json:"method"`
}

func NewBaseCommand(method string) BaseCommand {
	return BaseCommand{
		ID:     "1",
		Method: method,
	}
}
