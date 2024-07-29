package interactive

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"io"
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

	ip.Print("Welcome to interactive mode, type help for more info.")
	ip.Print("\n")
	ip.Print(PROMPT)

	ip.HandleCommand()
}

func (ip *InteractiveProgram) HandleCommand() {
	for {
		ip.resetInput()
		ip.scanner.Scan()
		line := ip.scanner.Text()

		ip.input = line
		ip.readChar()

		parsedLine := ip.parseLine()
		for _, item := range parsedLine {
			fmt.Println("item is ", item)
		}
		ip.Print("input line ", line, "\n")
		ip.Print(PROMPT)
	}
}

func (ip *InteractiveProgram) NextToken() Command {
	ip.skipWhitespace()
	cmd := Command{}

	switch ip.ch {
	case 0:
		cmd.Type = EOF
		cmd.Literal = "EOF"
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

func (ip *InteractiveProgram) parseLine() []Command {
	var cmds = []Command{}
	for ip.next <= len(ip.input) {
		cmd := ip.NextToken()
		cmds = append(cmds, cmd)
	}
	return cmds
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

func (ip *InteractiveProgram) resetInput() {
	ip.input = ""
	ip.cur = 0
	ip.next = 0
}
