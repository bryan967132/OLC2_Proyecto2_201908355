package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Cast struct {
	Line    int
	Column  int
	Destiny utils.Type
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewCast(line, column int, destiny utils.Type, exp interfaces.Expression) *Cast {
	return &Cast{line, column, destiny, utils.CAST, exp}
}

func (ct *Cast) LineN() int {
	return ct.Line
}

func (ct *Cast) ColumnN() int {
	return ct.Column
}

func (ct *Cast) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
