package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Pbasnal/linearequations/equation"

	"html/template"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {

	tpl = loadHtmlPages()
	setupTheRestAPI()
	//testEquationBuilder()
}

func loadHtmlPages() *template.Template {

	return template.Must(template.ParseGlob("templates/*.html"))
}

func setupTheRestAPI() {

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/eq", SolveEquations).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Index - sends the index page
func Index(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "index.html", nil)
}

func SolveEquations(w http.ResponseWriter, r *http.Request) {

	fmt.Println("PostMethod")
	r.ParseForm()

	fmt.Println("SolveEquations ", r.Form)
	equations := []equation.Equation{}

	var er error
	var solution map[string]float64

	for _, v := range r.Form {
		if len(v) == 0 || strings.TrimSpace(v[0]) == "" {
			er = errors.New("Empty equation provided")
			break
		}
		//fmt.Println(v[0])
		equations = append(equations, equation.BuildEquationFromText(v[0]))
	}

	if er == nil {
		solution, er = equation.SolveEquations(equations)
	}

	if er != nil {
		json.NewEncoder(w).Encode(er.Error())
		return
	}

	json.NewEncoder(w).Encode(solution)
}

func testEquationBuilder() {
	e1 := equation.BuildEquationFromText("0.8x +2.6y- 3 =3")
	e1.PrintNormalizedEquation()
	e1 = equation.BuildEquationFromText("1z+0y+8=3")
	e1.PrintNormalizedEquation()
	e1 = equation.BuildEquationFromText("7 y-3=3+7x")
	e1.PrintNormalizedEquation()
	e1 = equation.BuildEquationFromText("7 y-3+4x - z+a=3")
	e1.PrintNormalizedEquation()
}

func testEquations2() {
	var e1, e2, e3, e4 equation.Equation
	e1 = equation.CreateEquation()
	e1.AppendVariableToEqation(2, "x")
	e1.AppendVariableToEqation(2, "y")
	e1.AppendVariableToEqation(2, "z")
	e1.AppendVariableToEqation(2, "a")
	e1.AppendEqualToEqation()
	e1.AppendConstantToEqation(3)

	e2 = equation.CreateEquation()
	e2.AppendVariableToEqation(3, "x")
	e2.AppendVariableToEqation(4, "y")
	e2.AppendVariableToEqation(2, "z")
	e2.AppendVariableToEqation(2, "a")
	e2.AppendEqualToEqation()
	e2.AppendConstantToEqation(6)

	e3 = equation.CreateEquation()
	e3.AppendVariableToEqation(6, "x")
	e3.AppendVariableToEqation(9, "y")
	e3.AppendVariableToEqation(4, "z")
	e2.AppendVariableToEqation(2, "a")
	e3.AppendEqualToEqation()
	e3.AppendConstantToEqation(10)

	e4 = equation.CreateEquation()
	e4.AppendVariableToEqation(2, "x")
	e4.AppendVariableToEqation(11, "y")
	e4.AppendVariableToEqation(14, "z")
	e4.AppendVariableToEqation(9, "a")
	e4.AppendEqualToEqation()
	e4.AppendConstantToEqation(22)

	fmt.Println("Initial E1")
	e1.PrintNormalizedEquation()
	fmt.Println("Initial E2")
	e2.PrintNormalizedEquation()
	fmt.Println("Initial E3")
	e3.PrintNormalizedEquation()
	fmt.Println("Initial E4")
	e4.PrintNormalizedEquation()

	equation.SolveEquations([]equation.Equation{e1, e2, e3, e4})
}

func testEquations() {
	var e1, e2 equation.Equation
	e1 = equation.CreateEquation()
	e1.AppendVariableToEqation(2, "x")
	e1.AppendVariableToEqation(2, "y")
	e1.AppendEqualToEqation()
	e1.AppendConstantToEqation(3)

	e2 = equation.CreateEquation()
	e2.AppendVariableToEqation(6, "x")
	e2.AppendVariableToEqation(1, "y")
	e2.AppendEqualToEqation()
	e2.AppendConstantToEqation(6)

	fmt.Println("\nE1 : ")
	e1.PrintNormalizedEquation()

	fmt.Println("\nE2 : ")
	e2.PrintNormalizedEquation()

	fmt.Println("\nE1 - E2 : ")
	result := equation.SubtractEquations(e1, e2)
	result.PrintNormalizedEquation()
}

func testTree() {

	root := equation.CreateEquationTree()

	root.AddVariable(2, "z")
	root.AddEqualTo()
	root.AddConstant(3)
	root.AddVariable(2, "x")
	root.AddVariable(2, "y")
	root.PrintEquation()
}
