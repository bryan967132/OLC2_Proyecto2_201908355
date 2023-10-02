package expressions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type Primitive struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Value   interface{}
	Type    utils.Type
}

func NewPrimitive(line, column int, value interface{}, typeD utils.Type) *Primitive {
	return &Primitive{line, column, utils.PRIMITIVE, value, typeD}
}

func (ar *Primitive) LineN() int {
	return ar.Line
}

func (ar *Primitive) ColumnN() int {
	return ar.Column
}

func (pr *Primitive) Exec(env *env.Env) *utils.ReturnValue {

	switch pr.Type {
	case utils.INT:
	case utils.FLOAT:
	}
	return nil
}
