package equation

import (
	"errors"
	"fmt"
)

// Equation does something
type Equation struct {
	equationNodes  []equationNode
	normalizedEq   RootNode
	equalToNodeLoc int
}

// CreateEquation creates a new instance of the code
func CreateEquation() Equation {
	e := Equation{}
	e.equationNodes = []equationNode{}
	e.normalizedEq = CreateEquationTree()
	return e
}

// CreateTestEquation does something
func CreateTestEquation() Equation {
	twox := variableNode{2, "x"}
	twoy := variableNode{2, "y"}
	three := constantNode{3}

	e := Equation{}
	e.equationNodes = append(e.equationNodes, twox)
	e.equationNodes = append(e.equationNodes, twoy)
	e.equationNodes = append(e.equationNodes, equalToNode{})
	e.equationNodes = append(e.equationNodes, three)

	return e
}

// AppendVariableToEqation appends variable
func (e *Equation) AppendVariableToEqation(coeff float64, variable string) {
	if coeff < 0.001 && coeff > -0.001 {
		return //coeff = 0.0
	}
	e.normalizedEq.AddVariable(coeff, variable)

	node := variableNode{coeff, variable}
	e.equationNodes = append(e.equationNodes, node)
}

// AppendConstantToEqation appends variable
func (e *Equation) AppendConstantToEqation(value float64) {
	if value < 0.001 && value > -0.001 {
		return //	value = 0.0
	}
	e.normalizedEq.AddConstant(value)

	node := constantNode{value}
	e.equationNodes = append(e.equationNodes, node)
}

// AppendEqualToEqation appends variable
func (e *Equation) AppendEqualToEqation() {
	e.normalizedEq.AddEqualTo()

	node := equalToNode{}
	e.equationNodes = append(e.equationNodes, node)
	e.equalToNodeLoc = len(e.equationNodes) - 1
}

// PrintEquation prints equation
func (e Equation) PrintEquation() {
	for i := 0; i < len(e.equationNodes); i++ {
		v, t := e.equationNodes[i].getValueAndType()
		if i != 0 && t != "e" && v > 0.0 && i != e.equalToNodeLoc+1 {
			fmt.Printf(" + ")
		}

		e.equationNodes[i].print()
		fmt.Printf("  ")
	}
}

// PrintNormalizedEquation prints normalized equation
func (e Equation) PrintNormalizedEquation() {
	e.normalizedEq.PrintEquation()
	fmt.Println()
}

// GetFirstVariable gets the first variable in the equation
func (e Equation) getFirstVariable() (variableNode, error) {
	for _, v := range e.normalizedEq.variables {
		if v.coeff != 0 {
			return *v, nil
		}
	}
	return variableNode{}, errors.New("Equation does not have any variables")
}

// DivideEquationBy - divides the equation by the given number and returns a new equation
func DivideEquationBy(e Equation, num float64) Equation {
	result := CreateEquation()

	for k, v := range e.normalizedEq.variables {
		//fmt.Printf("\nVar: %s : %.2f", k, v.coeff)
		result.AppendVariableToEqation(v.coeff/num, k)
	}

	result.AppendEqualToEqation()
	result.AppendConstantToEqation(e.normalizedEq.constant.value / num)

	return result
}

// MultiplyEquationBy - divides the equation by the given number and returns a new equation
func MultiplyEquationBy(e Equation, num float64) Equation {
	result := CreateEquation()

	for k, v := range e.normalizedEq.variables {
		//fmt.Printf("\nVar: %s : %.2f", k, v.coeff)
		result.AppendVariableToEqation(v.coeff*num, k)
	}

	result.AppendEqualToEqation()
	result.AppendConstantToEqation(e.normalizedEq.constant.value * num)

	return result
}

// SubtractEquations - returns a new equation which is differnce of the two input equations
func SubtractEquations(e1 Equation, e2 Equation) Equation {
	result := CreateEquation()

	keys := []string{}
	for k, v := range e1.normalizedEq.variables {
		//fmt.Printf("\nVar: %s : %.2f", k, v.coeff)
		num := 0.0
		if _, ok := e2.normalizedEq.variables[k]; ok {
			num = e2.normalizedEq.variables[k].coeff
		}
		keys = append(keys, k)
		result.AppendVariableToEqation(v.coeff-num, k)
	}

	for k, v := range e2.normalizedEq.variables {
		//fmt.Printf("\nVar: %s : %.2f", k, v.coeff)

		if sliceContains(keys, k) {
			continue
		}
		num := -1 * v.coeff
		result.AppendVariableToEqation(num, k)
	}

	result.AppendEqualToEqation()
	result.AppendConstantToEqation(
		e1.normalizedEq.constant.value - e2.normalizedEq.constant.value)

	return result
}

func sliceContains(slice []string, val string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return true
		}
	}
	return false
}

func (e Equation) getCoeffOfVariable(variable string) (float64, error) {

	if v, ok := e.normalizedEq.variables[variable]; ok {
		return v.coeff, nil
	}

	return 0, errors.New("Variable does not exists")
}

func (e Equation) solve(solutions *map[string]float64) {

	for k, v := range *solutions {
		if _, ok := e.normalizedEq.variables[k]; ok {
			e.normalizedEq.constant.value -= e.normalizedEq.variables[k].coeff * v
			delete(e.normalizedEq.variables, k)
		}
	}

	if len(e.normalizedEq.variables) == 1 {
		variable, _ := e.getFirstVariable()
		(*solutions)[variable.variable] = e.normalizedEq.constant.value / variable.coeff
	}
}
