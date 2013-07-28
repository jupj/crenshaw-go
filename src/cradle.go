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

// Initialize
func Init() {
	GetChar()
}

// Main Program
func main() {
	Init()
}
