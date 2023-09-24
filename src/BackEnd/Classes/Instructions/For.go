package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type For struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	IDIter   string
	Exp      interfaces.Expression
	LimInf   interfaces.Expression
	LimSup   interfaces.Expression
	Block    interfaces.Instruction
}

func NewFor(line, column int, idIte string, exp, limInf, limSup interfaces.Expression, block interfaces.Instruction) *For {
	return &For{line, column, utils.LOOP_FOR, idIte, exp, limInf, limSup, block}
}

func (f *For) LineN() int {
	return f.Line
}

func (f *For) ColumnN() int {
	return f.Column
}

func (f *For) Exec(Env *env.Env) *utils.ReturnType {
	return nil
}
