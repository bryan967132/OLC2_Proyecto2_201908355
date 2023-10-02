package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type InitID struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       utils.Type
	Value      interfaces.Expression
}

func NewInitID(line, column int, isVariable bool, id string, Type utils.Type, value interfaces.Expression) *InitID {
	return &InitID{line, column, isVariable, utils.INIT_ID, id, Type, value}
}

func (in *InitID) LineN() int {
	return in.Line
}

func (in *InitID) ColumnN() int {
	return in.Column
}

func (in *InitID) Exec(env *env.Env) *utils.ReturnValue {
	return nil
}
