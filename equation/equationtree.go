package equation

import "fmt"

// RootNode defines the root of the tree
type RootNode struct {
	variables    map[string]*variableNode
	symbols      []string
	constant     constantNode
	equalToAdded bool
}

// CreateEquationTree - creates a new root node
func CreateEquationTree() RootNode {

	return RootNode{
		make(map[string]*variableNode),
		[]string{},
		constantNode{0.0},
		false,
	}
}

// AddVariable - adds a variable to the node
func (t *RootNode) AddVariable(coeff float64, variable string) {

	if t.equalToAdded {
		coeff *= -1.0
	}

	if _, ok := t.variables[variable]; ok {
		(*t.variables[variable]).add(coeff)
		return
	}

	node := variableNode{coeff, variable}
	t.variables[variable] = &node
	t.symbols = append(t.symbols, variable)
}

// AddConstant - adds a constant to the node
func (t *RootNode) AddConstant(value float64) {

	if !t.equalToAdded {
		value *= -1.0
	}
	t.constant.value += value
}

// AddEqualTo - sets the boolean which signifies that the equalto symbol has been passed
func (t *RootNode) AddEqualTo() {

	t.equalToAdded = true
}

// PrintEquation - prints the entire tree
func (t *RootNode) PrintEquation() {

	i := 0
	for _, v := range t.variables {
		if i != 0 && v.coeff > 0 {
			fmt.Printf(" + ")
		}
		fmt.Printf(" %.2f%s ", v.coeff, v.variable)
		i++
	}

	fmt.Printf(" = ")
	fmt.Printf(" %.2f ", t.constant.value)
}
