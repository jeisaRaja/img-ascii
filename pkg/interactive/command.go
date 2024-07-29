package interactive 

const (
	EXIT    = "EXIT"
	HELP    = "HELP"
	CONVERT = "CONVERT"
	VAL     = "VAL"
)

type CommandType string

type Command struct {
	Type    CommandType
	Literal string
}

var commands = map[string]CommandType{
	EXIT:    "exit",
	HELP:    "help",
	CONVERT: "convert",
}

func lookUpCommand(cmd string) CommandType {
	if command, ok := commands[cmd]; ok {
		return command
	}
	return VAL
}
