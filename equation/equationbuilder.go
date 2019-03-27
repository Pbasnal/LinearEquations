package equation

import (
	"strconv"
	"unicode"
)

// BuildEquationFromText - builds a linear equation from the given equaion in string format
func BuildEquationFromText(eq string) Equation {

	lenOfEq := len(eq)
	queue := []rune{}
	queueIndex := -1
	isEqualToAdded := false
	variablePart := -1

	var e1 Equation
	e1 = CreateEquation()

	for i := 0; i < lenOfEq; i++ {
		if eq[i] == ' ' {
			continue
		}
		if eq[i] == '+' || eq[i] == '-' {
			if queueIndex == -1 {
				continue
			}

			if queue[queueIndex] != '~' && queue[queueIndex] != '-' {
				addVariableOrConstToEquation(&queue, &queueIndex, &variablePart, &e1)
				if eq[i] == '-' {
					queue = append(queue, rune(eq[i]))
					queueIndex++
				}
			}
		} else if eq[i] == '=' {
			if isEqualToAdded {
				continue
			}
			addVariableOrConstToEquation(&queue, &queueIndex, &variablePart, &e1)
			e1.AppendEqualToEqation()
		} else {
			queue = append(queue, rune(eq[i]))
			queueIndex++
		}

		if unicode.IsLetter(rune(eq[i])) && variablePart == -1 {
			variablePart = queueIndex
		}
	}

	addVariableOrConstToEquation(&queue, &queueIndex, &variablePart, &e1)

	return e1
}

func addVariableOrConstToEquation(queue *[]rune, queueIndex *int, variablePart *int, e *Equation) {

	if unicode.IsLetter((*queue)[*queueIndex]) {
		num := 1.0
		var variable string
		if *variablePart == 0 {

		} else if unicode.IsNumber((*queue)[*variablePart-1]) {
			num, _ = strconv.ParseFloat(string((*queue)[0:*variablePart]), 64)
		} else if (*queue)[*variablePart-1] == '-' {
			num = -1.0
		}

		variable = string((*queue)[*variablePart:])

		e.AppendVariableToEqation(num, variable)
	} else if unicode.IsNumber((*queue)[*queueIndex]) {
		num, _ := strconv.ParseFloat(string((*queue)), 64)
		e.AppendConstantToEqation(num)
	}

	*queue = []rune{}
	*queueIndex = -1
	*variablePart = -1
}
