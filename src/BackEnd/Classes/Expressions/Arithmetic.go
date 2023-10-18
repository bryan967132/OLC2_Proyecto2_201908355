package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Arithmetic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
	Type    utils.Type
}

func NewArithmetic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Arithmetic {
	return &Arithmetic{line, column, utils.ARITHMETIC_OP, exp1, sign, exp2, utils.NIL}
}

func (ar *Arithmetic) LineN() int {
	return ar.Line
}

func (ar *Arithmetic) ColumnN() int {
	return ar.Column
}

func (ar *Arithmetic) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}