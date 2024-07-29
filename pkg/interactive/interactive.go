package interactive

import (
	"bufio"
	"errors"
	"image"
	"io"
	"strings"

	"github.com/jeisaraja/img-ascii/pkg/ascii"
)

var ErrExit = errors.New("exit")

type InteractiveProgram struct {
	scanner *bufio.Scanner
	input   string
	out     io.Writer
	img     image.Image
	cur     int
	next    int
	ch      byte
}

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	ip := &InteractiveProgram{scanner: scanner, out: out}
	ip.readChar()

	ip.Print("Welcome to interactive mode, type help for more info.")
	ip.Print("\n")
	for {
		ip.Print(PROMPT)
		err := ip.HandleCommand()
		if err != nil {
			break
		}
	}
}

func (ip *InteractiveProgram) HandleCommand() error {
	line := ip.parseREPL()
	cmd := lookUpCommand()
	switch strings.ToLower(line) {
	case EXIT:
		ip.Print("exiting img-ascii", "\n")
		return ErrExit
	case HELP:
		ip.Print("this is help", "\n")
		return nil
	case CONVERT:
		ascii.ImageToASCII()
		return nil
	default:
		return nil
	}
}

func (ip *InteractiveProgram) NextCommand() Command {
	ip.skipWhitespace()
	cmd := Command{}

	switch ip.ch {
	case 0:
	default:
		if isLetter(ip.ch) {
			cmd.Literal = ip.readCommand()
			cmd.Type = lookUpCommand(cmd.Literal)
			return cmd
		}
	}
	return cmd
}

func (ip *InteractiveProgram) Print(lines ...string) {
	for _, line := range lines {
		io.WriteString(ip.out, line)
	}
}

func (ip *InteractiveProgram) parseREPL() string {
	ip.scanner.Scan()
	line := ip.scanner.Text()
	return line
}

func (ip *InteractiveProgram) skipWhitespace() {
	for ip.ch == ' ' || ip.ch == '\t' || ip.ch == '\n' || ip.ch == '\r' {
		ip.readChar()
	}
}

func (ip *InteractiveProgram) readChar() {
	if ip.next >= len(ip.input) {
		ip.ch = 0
	} else {
		ip.ch = ip.input[ip.next]
	}
	ip.cur = ip.next
	ip.next++
}

func (ip *InteractiveProgram) readCommand() string {
	pos := ip.cur
	for isLetter(ip.ch) {
		ip.readChar()
	}

	return ip.input[pos:ip.cur]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
