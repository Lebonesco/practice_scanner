package scanner 

import (
	"github.com/Lebonesco/scanner/lexer"
	"github.com/Lebonesco/scanner/token"
	"testing"
)

// test that the lexer works
func TestParseQuotes(t *testing.T) {
	input := `Sed ut perspiciatis, "unde omnis" iste 'natus' en 
		"hello" "es`;

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.TokMap.Type("text"), "Sed"},
		{token.TokMap.Type("text"), "ut"},
		{token.TokMap.Type("text"), "perspiciatis,"},
		{token.TokMap.Type("doubleparen"), "\"unde omnis\""},
		{token.TokMap.Type("text"), "iste"},
		{token.TokMap.Type("singleparen"), "'natus'"},
		{token.TokMap.Type("text"), "en"},
		{token.TokMap.Type("doubleparen"), "\"hello\""},
		{token.TokMap.Type("text"), "\"es"},
	}

	l := lexer.NewLexer([]byte(input))

	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d", 
				i, tt.expectedType, tok.Type)
		}

		if string(tok.Lit) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", 
				i, tt.expectedLiteral, string(tok.Lit))
		}
	}
}

// test that quotation replacement works
func TestReplaceQuotes(t *testing.T) {
	tests := []struct {
		input string
		result string
	}{
		{
			input: `Sed ut perspiciatis, "unde omnis" iste 'natus' en 
			"hello" "es`,
			result: `Sed ut perspiciatis, &ldquounde omnis&rdquo iste &lquonatus&rquo en &ldquohello&rdquo "es`,
		},
	}

	for i, test := range tests {
		l := lexer.NewLexer([]byte(test.input))
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
		tmp = tmp[:len(tmp)-1] // remove extra end space
		if tmp != test.result {
			t.Fatalf("tests[%d] - [%s] doesn't match [%s]", i, tmp, test.result)
		}
	}
}