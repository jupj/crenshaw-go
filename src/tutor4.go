// Program Cradle
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var Look rune            // Lookahead Character
var Table map[string]int // Variable table
var inputReader = bufio.NewReader(os.Stdin)

// Read New Character From Input Stream
func GetChar() {
	Look, _, _ = inputReader.ReadRune()
}

// Report an Error
func Error(s string) {
	fmt.Println()
	fmt.Printf("\aError: %s.\n", s)
}

// Report Error and Halt
func Abort(s string) {
	Error(s)
	os.Exit(1)
}

// Report What Was Expected
func Expected(s string) {
	Abort(s + " Expected")
}

// Match a Specific Input Character
func Match(x rune) {
	if Look == x {
		GetChar()
	} else {
		Expected(fmt.Sprintf("%q", x))
	}
}

// Recognize and Skip Over a Newline
func NewLine() {
	if Look == '\r' {
		GetChar()
		if Look == '\n' {
			GetChar()
		}
	}
}

// Recognize an Alpha Character
func IsAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

// Recognize a Decimal Digit
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// Recognize an Alphanumeric
func IsAlNum(r rune) bool {
	return IsAlpha(r) || IsDigit(r)
}

// Recognize and Addop
func IsAddop(r rune) bool {
	return (r == '+') || (r == '-')
}

// Get an Identifier
func GetName() (result string) {
	if !IsAlpha(Look) {
		Expected("Name")
	}
	result = string(unicode.ToUpper(Look))
	GetChar()
	return
}

// Get a Number
func GetNum() int {
	Value := 0
	if !IsDigit(Look) {
		Expected("Integer")
	}
	for IsDigit(Look) {
		digit, _ := strconv.Atoi(string(Look))
		Value = 10*Value + digit
		GetChar()
	}
	return Value
}

// Output a String with Tab
func Emit(s string) {
	fmt.Print("\t", s)
}

// Output a String with Tab and CRLF
func EmitLn(s string) {
	Emit(s)
	fmt.Println()
}

// Parse and Translate a Math Factor
func Factor() (result int) {
	if Look == '(' {
		Match('(')
		result = Expression()
		Match(')')
	} else if IsAlpha(Look) {
		result = Table[GetName()]
	} else {
		result = GetNum()
	}
	return
}

// Parse and Translate a Math Term
func Term() int {
	Value := Factor()
	for (Look == '*') || (Look == '/') {
		switch Look {
		case '*':
			{
				Match('*')
				Value *= Factor()
			}
		case '/':
			{
				Match('/')
				Value /= Factor()
			}
		}
	}
	return Value
}

// Parse and Translate an Expression
func Expression() int {
	var Value int
	if IsAddop(Look) {
		Value = 0
	} else {
		Value = Term()
	}
	for IsAddop(Look) {
		switch Look {
		case '+':
			{
				Match('+')
				Value += Term()
			}
		case '-':
			{
				Match('-')
				Value -= Term()
			}
		}
	}
	return Value
}

// Parse and Translate an Assignment Statement
func Assignment() {
	Name := GetName()
	Match('=')
	Table[Name] = Expression()
}

// Input Routine
func Input() {
	Match('?')
    name := GetName()
    input := GetNum()
	Table[name] = input
}

// Output Routine
func Output() {
	Match('!')
	fmt.Println(Table[GetName()])
}

// Initialize the Variable Area
func InitTable() {
	// The default value of an non-existing int value is 0,
	// so we don't need to initialize the individual variables
	Table = make(map[string]int)
}

// Initialize
func Init() {
	InitTable()
	GetChar()
}

// Main Program
func main() {
	Init()
	for Look != '.' {
		switch Look {
		case '?':
			Input()
		case '!':
			Output()
		default:
			Assignment()
		}
		NewLine()
	}
}
