package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type If struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
	Except    interfaces.Instruction
}

func NewIf(line, column int, condition interfaces.Expression, block interfaces.Instruction, except interfaces.Instruction) *If {
	return &If{Line: line, Column: column, TypeInst: utils.IF, Condition: condition, Block: block, Except: except}
}

func (i *If) LineN() int {
	return i.Line
}

func (i *If) ColumnN() int {
	return i.Column
}

func (i *If) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
