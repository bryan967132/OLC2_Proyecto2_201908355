package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type AsignPosArray struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Index    []interfaces.Expression
	NewValue interfaces.Expression
}

func NewAsignPosArray(line, column int, id string, index []interfaces.Expression, newValue interfaces.Expression) *AsignPosArray {
	return &AsignPosArray{line, column, utils.ASIGN_ARRAY, id, index, newValue}
}

func (ac *AsignPosArray) LineN() int {
	return ac.Line
}

func (ac *AsignPosArray) ColumnN() int {
	return ac.Column
}

func (ac *AsignPosArray) Exec(env *env.Env) *utils.ReturnType {
	return nil
}
