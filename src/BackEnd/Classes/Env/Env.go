package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Env struct {
	Ids       *map[string]*Symbol
	Functions *map[string]*interface{}
	previous  *Env
	Name      string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{&map[string]*Symbol{}, &map[string]*interface{}{}, previous, name}
}

func (env *Env) SaveID(isVariable bool, id string, value *utils.ReturnType, Type utils.Type, line, column int) bool {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{IsVariable: isVariable, IsPrimitive: true, Value: value, Id: id, Type: Type}
		SymTable.Push(NewSymTab(line, column+1, isVariable, true, id, env.Name, Type, utils.NIL))
		return true
	}
	env.SetError("Redeclaración de variable existente", line, column)
	return false
}

func (env *Env) SaveArray(isVariable bool, id string, value interface{}, Type utils.Type, line, column int) bool {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{IsVariable: isVariable, IsPrimitive: false, Value: value, Id: id, Type: utils.VECTOR, ArrType: Type}
		SymTable.Push(NewSymTab(line, column+1, isVariable, false, id, env.Name, utils.VECTOR, Type))
		return true
	}
	env.SetError("Redeclaración de variable existente", line, column)
	return false
}

func (env *Env) GetValueID(id string, line, column int) *Symbol {
	var current *Env = env
	for current != nil {
		if _, exists := (*current.Ids)[id]; exists {
			return (*current.Ids)[id]
		}
		current = current.previous
	}
	current.SetError(fmt.Sprintf("Acceso a variable inexistente. '%s'", id), line, column)
	return nil
}

func (env *Env) ReasignID(id string, value *utils.ReturnType, line, column int) bool {
	var current *Env = env
	for current != nil {
		if _, exists := (*current.Ids)[id]; exists {
			if (*current.Ids)[id].IsVariable {
				if (*current.Ids)[id].Type == value.Type || (*current.Ids)[id].Type == utils.STRING && value.Type == utils.CHAR || (*current.Ids)[id].Type == utils.FLOAT && value.Type == utils.INT {
					(*current.Ids)[id].Value = value
					return true
				}
				current.SetError(fmt.Sprintf("Los tipos no coinciden en la asignación. Intenta asignar un \"%v\" a un \"%v\"", current.getType(value.Type), current.getType((*current.Ids)[id].Type)), line, column)
				return false
			}
			current.SetError("Resignación de valor a constante", line, column)
			return false
		}
		current = current.previous
	}
	current.SetError("Resignación de valor a variable inexistente", line, column)
	return false
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

func (env *Env) SetPrints(print string) {
	utils.PrintConsole = append(utils.PrintConsole, print)
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

func (env *Env) getType(Type utils.Type) string {
	switch Type {
	case utils.INT:
		return "Int"
	case utils.FLOAT:
		return "Float"
	case utils.STRING:
		return "String"
	case utils.BOOLEAN:
		return "Bool"
	case utils.CHAR:
		return "Character"
	default:
		return "nil"
	}
}
