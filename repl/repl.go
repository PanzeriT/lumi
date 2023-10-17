package repl

import (
	"bufio"
	"fmt"
	"github.com/panzerit/lumi/lexer"
	"github.com/panzerit/lumi/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%-10s%s\n", tok.Type+":", tok.Literal)

		}
	}
}
