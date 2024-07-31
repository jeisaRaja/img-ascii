package interactive

const (
	EXIT    = "EXIT"
	HELP    = "HELP"
	CONVERT = "CONVERT"
	VAL     = "VAL"
	EOF     = "EOF"
	INVALID = "INVALID"
)

type CommandType string

type Command struct {
	Type    CommandType
	Literal string
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
	return INVALID
}
