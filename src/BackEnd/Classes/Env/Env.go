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
