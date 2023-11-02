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
	c3dgen.AddComment("----------- If ------------")
	condition := i.Condition.Exec(Env, c3dgen)
	if condition.Type == utils.BOOLEAN {
		newLabel := c3dgen.NewLabel()
		for _, lbl := range condition.TrueLabel {
			c3dgen.AddLabel(lbl)
		}
		block := i.Block.Exec(Env, c3dgen)

		block.OutLabel = append(block.OutLabel, newLabel)
		for _, lbl := range block.OutLabel {
			c3dgen.AddGoto(lbl)
		}

		for _, lbl := range condition.FalseLabel {
			c3dgen.AddLabel(lbl)
		}

		// else
		if i.Except != nil {
			i.Except.Exec(Env, c3dgen)
		}

		for _, lbl := range block.OutLabel {
			c3dgen.AddLabel(lbl)
		}
		c3dgen.AddComment("---------------------------")
		return nil
	}
	c3dgen.AddComment("---------------------------")
	Env.SetError("No se evalúa una expresión lógica o relacional como condicion", i.Line, i.Column)
	return nil
}
