package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type While struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
}

func NewWhile(line, column int, condition interfaces.Expression, block interfaces.Instruction) *While {
	return &While{line, column, utils.LOOP_WHILE, condition, block}
}

func (w *While) LineN() int {
	return w.Line
}

func (w *While) ColumnN() int {
	return w.Column
}

func (w *While) Exec(Env *env.Env) *utils.ReturnType {
	return nil
}
