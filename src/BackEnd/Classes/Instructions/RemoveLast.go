package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type RemoveLast struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
}

func NewRemoveLast(line, column int, id string) *RemoveLast {
	return &RemoveLast{line, column, utils.ARRAY_REMOVELAST, id}
}

func (r *RemoveLast) LineN() int {
	return r.Line
}

func (r *RemoveLast) ColumnN() int {
	return r.Column
}

func (r *RemoveLast) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
