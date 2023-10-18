package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Logic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewLogic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Logic {
	return &Logic{line, column, utils.LOGIC_OP, exp1, sign, exp2}
}

func (lg *Logic) LineN() int {
	return lg.Line
}

func (lg *Logic) ColumnN() int {
	return lg.Column
}

func (lg *Logic) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
