package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type AccessID struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Id      string
}

func NewAccessID(line, column int, id string) *AccessID {
	return &AccessID{line, column, utils.ACCESS_ID, id}
}

func (ac *AccessID) LineN() int {
	return ac.Line
}

func (ac *AccessID) ColumnN() int {
	return ac.Column
}

func (ac *AccessID) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
