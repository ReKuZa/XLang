package main

import(
	"fmt"
)

type token struct {
	value string
	kind string
	line int
}
type custom struct {
	kind string
	value string
}
var fn = custom{
	kind : "function",
	value : "block",
}
var variable = custom{
	kind : "variable",
	value : "var",
}
var integer = custom{
	kind : "integer",
	value : "int",
}
var uinteger = custom{
	kind : "uinteger",
	value : "uint",
}
var boolean = custom{
	kind : "boolean",
	value : "bool",
}
var lot = []custom{fn,variable,integer,uinteger,boolean}
func Tokenizer(input string) []token  {
	input += "\n"
	current := 0
	line := 0
	tokens := []token{}
	for current <  len([]rune(input)) {
		char := string([]rune(input)[current])
		if char == "(" {
			tokens = append(tokens, token{
				kind : "oparen",
				value : "(",
			})

			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{
				kind : "cparen",
				value : ")",
			})
		}
		if char == " " {
			current++
			continue
		}
		if char == "\n" {
			current++
			line++
			continue
		}
		if isNumber(char) {
			value := ""
			for isNumber(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}
			
			tokens = append(tokens, token{
				kind : "number",
				value : value,
				line : line,
			})

			continue
		}
		if isLetter(char) {
			value := ""

			for isLetter(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}
			var isName = true
			for i := range lot {
				var val = lot[i].value
				var kd = lot[i].kind
				if value == val {
					tokens = append(tokens, token{
						kind : kd,
						value : value,
						line : line,
					})
					isName = false
					continue
				}
			}
			if isName != false {
				tokens = append(tokens,token{
					kind : "name",
					value : value,
					line : line,
				})
			}
			continue
		} else {
			current++
			continue
		}
		break
	}
	return tokens
}
func isNumber(char string) bool  {
	if char == "" {
		 return false
    }

    n := []rune(char)[0]

    if n >= '0' && n <= '9' {
    	return true
    }

    return false
}
func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	
	if n >= 'a' && n <= 'z' {
		return true
	}

	return false
}
func Compile(input string) string {
	getlex := Tokenizer(input)
	fmt.Println(getlex)
	return "ok"
}

func main() {
	j := `abc
	8383
	bye
	ok`
	a := Tokenizer(j)
	fmt.Println(a)
}
