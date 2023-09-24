package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Guard struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
}

func NewGuard(line, column int, condition interfaces.Expression, block interfaces.Instruction) *Guard {
	return &Guard{line, column, utils.GUARD, condition, block}
}

func (g *Guard) LineN() int {
	return g.Line
}

func (g *Guard) ColumnN() int {
	return g.Column
}

func (g *Guard) Exec(Env *env.Env) *utils.ReturnType {
	return nil
}
