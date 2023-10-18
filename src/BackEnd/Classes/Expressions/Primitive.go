package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
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

func (pr *Primitive) LineN() int {
	return pr.Line
}

func (pr *Primitive) ColumnN() int {
	return pr.Column
}

func (pr *Primitive) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	switch pr.Type {
	case utils.INT:
	case utils.FLOAT:
	case utils.BOOLEAN:
	case utils.NIL:
	default:
	}
	return nil
}
