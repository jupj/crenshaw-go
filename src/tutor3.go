// Program Cradle
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var Look rune // Lookahead Character
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

// Recognize an Alpha Character
func IsAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

// Recognize a Decimal Digit
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
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
func GetNum() (result string) {
	if !IsDigit(Look) {
		Expected("Integer")
	}
	result = string(Look)
	GetChar()
	return
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

// Parse and Translate an Identifier
func Ident() {
	Name := GetName()
	if Look == '(' {
		Match('(')
		Match(')')
		EmitLn("BSR " + Name)
	} else {
		EmitLn("MOVE " + Name + "(PC),D0")
	}
}

// Parse and Translate a Math Factor
func Factor() {
	if Look == '(' {
		Match('(')
		Expression()
		Match(')')
	} else if IsAlpha(Look) {
		Ident()
	} else {
		EmitLn("MOVE #" + GetNum() + ",D0")
	}
}

// Recognize and Translate a Multiply
func Multiply() {
	Match('*')
	Factor()
	EmitLn("MULS (SP)+,D0")
}

// Recognize and Translate a Divide
func Divide() {
	Match('/')
	Factor()
	EmitLn("MOVE (SP)+,D1")
	EmitLn("DIVS D1,D0")
}

// Parse and Translate a Math Term
func Term() {
	Factor()
	for (Look == '*') || (Look == '/') {
		EmitLn("MOVE D0,-(SP)")
		switch Look {
		case '*':
			Multiply()
		case '/':
			Divide()
		}
	}
}

// Recognize and Translate an Add
func Add() {
	Match('+')
	Term()
	EmitLn("ADD (SP)+,D0")
}

// Recognize and Translate a Subtract
func Subtract() {
	Match('-')
	Term()
	EmitLn("SUB (SP)+,D0")
	EmitLn("NEG D0")
}

// Parse and Translate an Expression
func Expression() {
	if IsAddop(Look) {
		EmitLn("CLR D0")
	} else {
		Term()
	}
	for (Look == '+') || (Look == '-') {
		EmitLn("MOVE D0,-(SP)")
		switch Look {
		case '+':
			Add()
		case '-':
			Subtract()
		}
	}
}

// Parse and Translate an Assignment Statement
func Assignment() {
	Name := GetName()
	Match('=')
	Expression()
	EmitLn("LEA " + Name + "(PC),A0")
	EmitLn("MOVE D0,(A0)")
}

// Initialize
func Init() {
	GetChar()
}

// Main Program
func main() {
	Init()
	Assignment()
	if Look != '\r' {
		Expected("Newline")
	}
}
