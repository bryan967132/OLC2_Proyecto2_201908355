package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Env struct {
	Ids         *map[string]*Symbol
	Functions   *map[string]*interface{}
	Size        int
	previous    *Env
	Name        string
	ContinueLbl string
	BreakLbl    []string
	ReturnLbl   string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{Ids: &map[string]*Symbol{}, Functions: &map[string]*interface{}{}, Size: 0, previous: previous, Name: name, ContinueLbl: "", BreakLbl: []string{}, ReturnLbl: ""}
}

func (env *Env) SaveID(isVariable bool, id string, value *utils.ReturnValue, Type utils.Type, line, column int) *Symbol {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{IsVariable: isVariable, IsPrimitive: true, Id: id, Type: Type, Position: env.Size, IsGlobal: env.Name == "Global"}
		SymTable.Push(NewSymTab(line, column+1, isVariable, true, id, env.Name, Type, utils.NIL))
		env.Size += 1
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

func (env *Env) SaveFunction(id string, Func *interface{}, Type utils.Type, line, column int) bool {
	if _, exists := (*env.Functions)[id]; !exists {
		(*env.Functions)[id] = Func
		SymTable.Push(NewSymTab(line, column+1, false, false, id, env.Name, Type, utils.NIL))
		return true
	}
	env.SetError("Redefinición de función existente", line, column)
	return false
}

func (env *Env) GetFunction(id string, line, column int) *interface{} {
	if _, exists := (*env.Functions)[id]; exists {
		return (*env.Functions)[id]
	}
	env.SetError("Acceso a función inexistente", line, column)
	return nil
}

func (env *Env) GetGlobal() *Env {
	current := env
	for current.previous != nil {
		current = current.previous
	}
	return current
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
