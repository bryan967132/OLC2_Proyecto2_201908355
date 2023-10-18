package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Relational struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewRelational(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Relational {
	return &Relational{line, column, utils.RELATIONAL_OP, exp1, sign, exp2}
}

func (rl *Relational) LineN() int {
	return rl.Line
}

func (rl *Relational) ColumnN() int {
	return rl.Column
}

func (rl *Relational) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
