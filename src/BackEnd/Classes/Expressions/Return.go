package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Return struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewReturn(line, column int, exp interfaces.Expression) *Return {
	return &Return{line, column, utils.RETURN, exp}
}

func (r *Return) LineN() int {
	return r.Line
}

func (r *Return) ColumnN() int {
	return r.Column
}

func (r *Return) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
