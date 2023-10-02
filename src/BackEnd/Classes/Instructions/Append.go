package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Append struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
	Exp      interfaces.Expression
}

func NewAppend(line, column int, id string, exp interfaces.Expression) *Append {
	return &Append{line, column, utils.ARRAY_APPEND, id, exp}
}

func (a *Append) LineN() int {
	return a.Line
}

func (a *Append) ColumnN() int {
	return a.Column
}

func (a *Append) Exec(env *env.Env) *utils.ReturnValue {
	return nil
}
