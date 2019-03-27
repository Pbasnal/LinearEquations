package equation

import "fmt"

// Equation holds the Raw Equation and normalized equations
type Equation struct {
	equationNodes  []equationNode
	equalToNodeLoc int

	normalizedEq RootNode
}

// equationNode holds one unit of the equation like '2x' or '5' or '='
type equationNode interface {
	print()
	getValueAndType() (float64, string)
}

type variableNode struct {
	coeff    float64
	variable string
}

func (v *variableNode) add(num float64) {
	v.coeff += num
}

func (v variableNode) print() {
	fmt.Printf("%.2f %s", v.coeff, v.variable)
}

func (v variableNode) getValueAndType() (float64, string) {
	return v.coeff, "v"
}

type constantNode struct {
	value float64
}

func (c *constantNode) add(num float64) {
	c.value += num
}

func (c constantNode) getValueAndType() (float64, string) {
	return c.value, "c"
}

func (c constantNode) print() {
	fmt.Printf("%.2f", c.value)
}

type equalToNode struct{}

func (eq equalToNode) getValueAndType() (float64, string) {
	return 0, "e"
}

func (eq equalToNode) print() {
	fmt.Printf("=")
}
