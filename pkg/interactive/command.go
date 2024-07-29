package interactive

const (
	EXIT    = "EXIT"
	HELP    = "HELP"
	CONVERT = "CONVERT"
	VAL     = "VAL"
	EOF     = "EOF"
)

type CommandType string

type Command struct {
	Type    CommandType
	Literal string
	Value   int
}

var commands = map[string]CommandType{
	"exit":    EXIT,
	"help":    HELP,
	"convert": CONVERT,
}

func lookUpCommand(cmd string) CommandType {
	if command, ok := commands[cmd]; ok {
		return command
	}
	return VAL
}
