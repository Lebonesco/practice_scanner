package main

import (
	"fmt"
	"github.com/Lebonesco/practice_scanner/lexer"
	"github.com/Lebonesco/practice_scanner/token"
	"io/ioutil"
	"os"
	"path/filepath"
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
