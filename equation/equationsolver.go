package equation

import (
	"errors"
	"fmt"
)

// SolveEquations takes a list of equations and solves them to get the value of the
// variables
func SolveEquations(equations []Equation) (map[string]float64, error) {
	variablesToSolve := getAllVariables(equations)
	if len(variablesToSolve) < len(equations) {
		fmt.Println("Insufficient information")
		return nil, errors.New("Insufficient information")
	}
	solution, er := elimination2(equations, variablesToSolve)

	if er != nil {
		fmt.Println(er)
		return nil, errors.New(er.Error())
	}

	// fmt.Println("Solution to the equations: ")
	// for k, v := range solution {
	// 	fmt.Printf("%s : %.2f\n", k, v)
	// }
	fmt.Println("Solved")
	return solution, nil
}

func swapEquations(equations *[]Equation, startIndex int, variable string) {
	maxCoeffofVar, err := (*equations)[startIndex].getCoeffOfVariable(variable)
	maxCoeffIndex := startIndex

	if err != nil {
		return
	}

	for i := startIndex + 1; i < len(*equations); i++ {
		coeff, err := (*equations)[i].getCoeffOfVariable(variable)
		if err != nil {
			continue
		}
		if maxCoeffofVar < coeff {
			maxCoeffIndex = i
			maxCoeffofVar = coeff
		}
	}
	if startIndex != maxCoeffIndex {
		tempE := (*equations)[startIndex]
		(*equations)[startIndex] = (*equations)[maxCoeffIndex]
		(*equations)[maxCoeffIndex] = tempE
	}
}

func getAllVariables(equations []Equation) []string {
	variablesToSolve := []string{}
	for i := 0; i < len(equations); i++ {
		for j := 0; j < len(equations[i].normalizedEq.symbols); j++ {
			sym := equations[i].normalizedEq.symbols[j]
			addSymbol := true
			for k := 0; k < len(variablesToSolve); k++ {
				if variablesToSolve[k] == sym {
					addSymbol = false
					break
				}
			}
			if addSymbol {
				variablesToSolve = append(variablesToSolve, sym)
			}
		}
	}
	return variablesToSolve
}

func elimination2(equations []Equation, variablesToSolve []string) (map[string]float64, error) {

	solutions := make(map[string]float64)
	for v := 0; v < len(variablesToSolve); v++ {
		for i := 0; i < len(equations); i++ {
			swapEquations(&equations, i, variablesToSolve[v])

			coeff, err := equations[i].getCoeffOfVariable(variablesToSolve[v])
			if err != nil {
				continue
			}
			equations[i] = DivideEquationBy(equations[i], coeff)

			backtrack, err := reduceEquations(i, &equations, variablesToSolve[v], &solutions)

			if err != nil {
				return solutions, err
			}

			if !backtrack {
				continue
			}
			for j := 0; j < len(equations); j++ {
				equations[j].solve(&solutions)
				//fmt.Printf("%d - ", j)
				//fmt.Println(solutions)
			}
		}
	}
	return solutions, nil
}

func reduceEquations(startIndex int, equations *[]Equation, variable string, solutions *map[string]float64) (bool, error) {
	backtrack := false
	for j := startIndex + 1; j < len(*equations); j++ {
		value, err := (*equations)[j].getCoeffOfVariable(variable)
		if err != nil {
			continue
		}
		e := MultiplyEquationBy((*equations)[startIndex], value)
		(*equations)[j] = SubtractEquations((*equations)[j], e)

		if len((*equations)[j].normalizedEq.variables) == 0 {
			return backtrack, errors.New("No Solution")
		}
		if len((*equations)[j].normalizedEq.variables) == 1 {
			(*equations)[j].solve(solutions)
			backtrack = true
		}
	}
	return backtrack, nil
}
