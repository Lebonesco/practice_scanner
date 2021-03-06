# practice_scanner

## Goal
practice writing a scanner with gocc 

## How to Run

### Download the code

First make sure that you have Golang installed on your system
and that you have your `$GOPATH` set. It is normally
defaulted to `$HOME/go`

Go is really particular when it comes
to import references having absolute paths.
That's why when you:
`git clone https://github.com/Lebonesco/practice_scanner.git`
make sure that it is in the file path:
`$GOPATH/src/github.com/Lebonesco/`

To run program with test file/file path:
`go run main.go text.txt`

To run unit tests:
```
go test -v
=== RUN   TestParseQuotes
--- PASS: TestParseQuotes (0.00s)
=== RUN   TestReplaceQuotes
--- PASS: TestReplaceQuotes (0.00s)
PASS
ok      github.com/Lebonesco/scanner    0.248s

```
