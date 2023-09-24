package expressions

import (
	env "TSwift/Classes/Env"
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

func (r *Return) Exec(env *env.Env) *utils.ReturnType {
	return &utils.ReturnType{Value: r.TypeExp, Type: utils.NIL}
}
