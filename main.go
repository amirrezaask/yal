package main

import "fmt"

type SymbolTable map[string]interface{}

type Sexp struct {
	Op   string
	Args []Value
}

const (
	_ = iota
	TypeString
)

type Value struct {
	Type string
	Data interface{}
}

const (
	StateNil = iota
	StateInList
	StateInArray
	StateInString
	StateInNumber
	StateInSymbol
)

func lastState(states []int) int {
	return states[len(states)-1]
}
func popState(states []int) {
	states = states[:len(states)-1]
}
func parse(code string) interface{} {
	i := 0
	var c byte
	var states []int
	var currentData interface{}
	// `"salam"`
	for {
		if states == nil {
			states = append(states, StateNil)
		}
		fmt.Println("current state is", lastState(states))
		if i > len(code)-1 {
			//EOF
			return currentData
		}
		c = code[i]
		switch lastState(states) {
		case StateNil:
			if c == '"' {
				states = append(states, StateInString)
				i++
				continue
			} else if c == '[' {
				states = append(states, StateInArray)
				i++
				continue
			} else if c == '(' {
				states = append(states, StateInList)
				i++
				continue
			} else if c >= '0' && c <= '9' {
				states = append(states, StateInNumber)
				i++
				continue
			} else if (c >= 'A' && c <= 'Z') || (c <= 'a' && c >= 'z') {
				states = append(states, StateInSymbol)
				i++
				continue
			}
		case StateInString:
			if currentData == nil {
				currentData = ""
			}
			if c == '"' {
				popState(states)
				i++
				continue
			}
			currentData = currentData.(string) + string(c)
			i++
			continue
		}
	}

}
func main() {
	fmt.Println(parse(`"Salam"`))
}
