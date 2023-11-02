package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type AccessArray struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Id      string
	Index   []interfaces.Expression
}

func NewAccessArray(line, column int, id string, index []interfaces.Expression) *AccessArray {
	return &AccessArray{line, column, utils.ACCESS_ARRAY, id, index}
}

func (ac *AccessArray) LineN() int {
	return ac.Line
}

func (ac *AccessArray) ColumnN() int {
	return ac.Column
}

func (ac *AccessArray) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
