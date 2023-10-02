package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type InitVector struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       *utils.AttribsType
	Value      *vector.Vector
}

func NewInitVector(line, column int, isVar bool, id string, Type *utils.AttribsType, value *vector.Vector) *InitVector {
	return &InitVector{line, column, isVar, utils.INIT_ID, id, Type, value}
}

func (in *InitVector) LineN() int {
	return in.Line
}

func (in *InitVector) ColumnN() int {
	return in.Column
}

func (in *InitVector) Exec(env *env.Env) *utils.ReturnValue {
	return nil
}
