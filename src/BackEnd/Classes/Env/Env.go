package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Env struct {
	Ids         *map[string]*Symbol
	Functions   *map[string]*interface{}
	Size        *map[string]int
	previous    *Env
	Name        string
	ContinueLbl []string
	BreakLbl    []string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{Ids: &map[string]*Symbol{}, Functions: &map[string]*interface{}{}, Size: &map[string]int{"size": 0}, previous: previous, Name: name, ContinueLbl: []string{}, BreakLbl: []string{}}
}

func (env *Env) SaveID(isVariable bool, id string, value *utils.ReturnValue, Type utils.Type, line, column int) *Symbol {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{IsVariable: isVariable, IsPrimitive: true, Id: id, Type: Type, Position: (*env.Size)["size"]}
		SymTable.Push(NewSymTab(line, column+1, isVariable, true, id, env.Name, Type, utils.NIL))
		(*env.Size)["size"] += 1
		return (*env.Ids)[id]
	}
	env.SetError("Redeclaración de variable existente", line, column)
	return nil
}

func (env *Env) GetValueID(id string, line, column int) *Symbol {
	var current *Env = env
	for current != nil {
		if symbol, exists := (*current.Ids)[id]; exists {
			return symbol
		}
		current = current.previous
	}
	current.SetError(fmt.Sprintf("Acceso a variable inexistente. '%s'", id), line, column)
	return nil
}

func (env *Env) PrintPrints() {
	fmt.Println("\nTSwift:")
	if len(utils.PrintConsole) > 0 {
		for _, Print := range utils.PrintConsole {
			fmt.Println(Print)
		}
	}
}

func (env *Env) SetError(errorD string, line, column int) {
	if !env.match(errorD, line, column+1) {
		utils.Errors = append(utils.Errors, *utils.NewError(line, column+1, utils.SEMANTIC, errorD))
	}
}

func (env *Env) match(err string, line, column int) bool {
	for _, s := range utils.Errors {
		if s.ToString() == (*utils.NewError(line, column, utils.SEMANTIC, err)).ToString() {
			return true
		}
	}
	return false
}

func (env *Env) PrintErrors() {
	if len(utils.Errors) > 0 {
		fmt.Println("\nERRORES:")
		for _, Error := range utils.Errors {
			fmt.Println(Error)
		}
	}
}
