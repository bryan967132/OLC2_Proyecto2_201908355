package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Remove struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
	Exp      interfaces.Expression
}

func NewRemove(line, column int, id string, exp interfaces.Expression) *Remove {
	return &Remove{line, column, utils.ARRAY_REMOVE, id, exp}
}

func (r *Remove) LineN() int {
	return r.Line
}

func (r *Remove) ColumnN() int {
	return r.Column
}

func (r *Remove) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
