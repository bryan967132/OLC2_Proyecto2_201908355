package instructions

import (
	env "TSwift/Classes/Env"
	expressions "TSwift/Classes/Expressions"
	C3DGen "TSwift/Classes/Generator"
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

func (g *Guard) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("---------- Guard ----------")
	condition := g.Condition.Exec(Env, c3dgen)
	if condition.Type == utils.BOOLEAN {
		if g.validateInstruction(Env) {
			newLabel := c3dgen.NewLabel()
			for _, lbl := range condition.FalseLabel {
				c3dgen.AddLabel(lbl)
			}

			block := g.Block.Exec(Env, c3dgen)

			block.OutLabel = append(block.OutLabel, newLabel)
			for _, lbl := range block.OutLabel {
				c3dgen.AddGoto(lbl)
			}

			for _, lbl := range condition.TrueLabel {
				c3dgen.AddLabel(lbl)
			}

			for _, lbl := range block.OutLabel {
				c3dgen.AddLabel(lbl)
			}
			c3dgen.AddComment("---------------------------")
			return nil
		}
		c3dgen.AddComment("---------------------------")
		Env.SetError("No se encuentra la instrucción de transferencia final", g.Line, g.Column)
		return nil
	}
	c3dgen.AddComment("---------------------------")
	Env.SetError("No se evalúa una expresión lógica o relacional como condicion", g.Line, g.Column)
	return nil
}

func (g *Guard) validateInstruction(env *env.Env) bool {
	instructions := g.Block.(*Block).Instructions
	if len(instructions) > 0 {
		if _, ok := instructions[len(instructions)-1].(*expressions.Return); ok {
			return true
		}
		if _, ok := instructions[len(instructions)-1].(*Continue); ok {
			return true
		}
		if _, ok := instructions[len(instructions)-1].(*Break); ok {
			return true
		}
	}
	return false
}
