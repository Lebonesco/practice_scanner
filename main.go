package main

import (
	"github.com/Lebonesco/scanner/lexer"
	"github.com/Lebonesco/scanner/token"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("no valid file name or path provided provided!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer([]byte(data))
	tok := l.Scan()
	var tmp string

	for tok.Type != token.TokMap.Type("$") { // keep scanning unitl end
		if tok.Type == token.TokMap.Type("doubleparen") {
			tmp += "&ldquo" + string(tok.Lit[1:len(tok.Lit)-1]) + "&rdquo"
		} else if tok.Type == token.TokMap.Type("singleparen") {
			tmp += "&lquo" + string(tok.Lit[1:len(tok.Lit)-1]) + "&rquo"
		} else {
			tmp += string(tok.Lit)
		}

		tmp += " "
		tok = l.Scan()
	}
	fmt.Printf("input: %s\n", data)
	fmt.Printf("result: %s", tmp)
}
