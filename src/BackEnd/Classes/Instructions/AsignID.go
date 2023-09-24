package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type AsignID struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Value    interfaces.Expression
}

func NewAsignID(line, column int, id string, value interfaces.Expression) *AsignID {
	return &AsignID{line, column, utils.ASIGN_ID, id, value}
}

func (as *AsignID) LineN() int {
	return as.Line
}

func (as *AsignID) ColumnN() int {
	return as.Column
}

func (as *AsignID) Exec(env *env.Env) *utils.ReturnType {
	return nil
}
