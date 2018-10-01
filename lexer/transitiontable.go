// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 4
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 2
		case r == 45: // ['-','-']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 5
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 6
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 2
		case r == 45: // ['-','-']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 7
		case r == 34: // ['"','"']
			return 8
		case r == 35: // ['#','#']
			return 7
		case r == 36: // ['$','$']
			return 7
		case r == 37: // ['%','%']
			return 7
		case r == 38: // ['&','&']
			return 7
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 7
		case r == 43: // ['+','+']
			return 7
		case r == 44: // [',',',']
			return 7
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 7
		case r == 61: // ['=','=']
			return 7
		case r == 63: // ['?','?']
			return 7
		case 65 <= r && r <= 90: // ['A','Z']
			return 7
		case r == 94: // ['^','^']
			return 7
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 7
		case r == 123: // ['{','{']
			return 7
		case r == 124: // ['|','|']
			return 7
		case r == 125: // ['}','}']
			return 7
		case r == 126: // ['~','~']
			return 7
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 7
		default:
			return 10
		}
	},
	// S4
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 11
		case r == 34: // ['"','"']
			return 8
		case r == 35: // ['#','#']
			return 11
		case r == 36: // ['$','$']
			return 11
		case r == 37: // ['%','%']
			return 11
		case r == 38: // ['&','&']
			return 11
		case r == 39: // [''',''']
			return 12
		case r == 42: // ['*','*']
			return 11
		case r == 43: // ['+','+']
			return 11
		case r == 44: // [',',',']
			return 11
		case r == 45: // ['-','-']
			return 11
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		case r == 61: // ['=','=']
			return 11
		case r == 63: // ['?','?']
			return 11
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 94: // ['^','^']
			return 11
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 123: // ['{','{']
			return 11
		case r == 124: // ['|','|']
			return 11
		case r == 125: // ['}','}']
			return 11
		case r == 126: // ['~','~']
			return 11
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 11
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 5
		case r == 39: // [''',''']
			return 6
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 5
		case r == 39: // [''',''']
			return 6
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 7
		case r == 34: // ['"','"']
			return 13
		case r == 35: // ['#','#']
			return 7
		case r == 36: // ['$','$']
			return 7
		case r == 37: // ['%','%']
			return 7
		case r == 38: // ['&','&']
			return 7
		case r == 39: // [''',''']
			return 6
		case r == 42: // ['*','*']
			return 7
		case r == 43: // ['+','+']
			return 7
		case r == 44: // [',',',']
			return 7
		case r == 45: // ['-','-']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 7
		case r == 61: // ['=','=']
			return 7
		case r == 63: // ['?','?']
			return 7
		case 65 <= r && r <= 90: // ['A','Z']
			return 7
		case r == 94: // ['^','^']
			return 7
		case r == 95: // ['_','_']
			return 7
		case 97 <= r && r <= 122: // ['a','z']
			return 7
		case r == 123: // ['{','{']
			return 7
		case r == 124: // ['|','|']
			return 7
		case r == 125: // ['}','}']
			return 7
		case r == 126: // ['~','~']
			return 7
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 7
		default:
			return 10
		}
	},
	// S8
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 8
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 2
		case r == 45: // ['-','-']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 8
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 2
		case r == 45: // ['-','-']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 34: // ['"','"']
			return 15
		case r == 35: // ['#','#']
			return 14
		case r == 36: // ['$','$']
			return 14
		case r == 37: // ['%','%']
			return 14
		case r == 38: // ['&','&']
			return 14
		case r == 42: // ['*','*']
			return 14
		case r == 43: // ['+','+']
			return 14
		case r == 44: // [',',',']
			return 14
		case r == 45: // ['-','-']
			return 14
		case 48 <= r && r <= 57: // ['0','9']
			return 14
		case r == 61: // ['=','=']
			return 14
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 94: // ['^','^']
			return 14
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		case r == 123: // ['{','{']
			return 14
		case r == 124: // ['|','|']
			return 14
		case r == 125: // ['}','}']
			return 14
		case r == 126: // ['~','~']
			return 14
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 14
		default:
			return 10
		}
	},
	// S11
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 11
		case r == 34: // ['"','"']
			return 5
		case r == 35: // ['#','#']
			return 11
		case r == 36: // ['$','$']
			return 11
		case r == 37: // ['%','%']
			return 11
		case r == 38: // ['&','&']
			return 11
		case r == 39: // [''',''']
			return 16
		case r == 42: // ['*','*']
			return 11
		case r == 43: // ['+','+']
			return 11
		case r == 44: // [',',',']
			return 11
		case r == 45: // ['-','-']
			return 11
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		case r == 61: // ['=','=']
			return 11
		case r == 63: // ['?','?']
			return 11
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 94: // ['^','^']
			return 11
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 123: // ['{','{']
			return 11
		case r == 124: // ['|','|']
			return 11
		case r == 125: // ['}','}']
			return 11
		case r == 126: // ['~','~']
			return 11
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 11
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 8
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 44: // [',',',']
			return 2
		case r == 45: // ['-','-']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 5
		case r == 39: // [''',''']
			return 6
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 34: // ['"','"']
			return 15
		case r == 35: // ['#','#']
			return 14
		case r == 36: // ['$','$']
			return 14
		case r == 37: // ['%','%']
			return 14
		case r == 38: // ['&','&']
			return 14
		case r == 42: // ['*','*']
			return 14
		case r == 43: // ['+','+']
			return 14
		case r == 44: // [',',',']
			return 14
		case r == 45: // ['-','-']
			return 14
		case 48 <= r && r <= 57: // ['0','9']
			return 14
		case r == 61: // ['=','=']
			return 14
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 94: // ['^','^']
			return 14
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		case r == 123: // ['{','{']
			return 14
		case r == 124: // ['|','|']
			return 14
		case r == 125: // ['}','}']
			return 14
		case r == 126: // ['~','~']
			return 14
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 14
		default:
			return 10
		}
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 5
		case r == 39: // [''',''']
			return 6
		}
		return NoState
	},
}
