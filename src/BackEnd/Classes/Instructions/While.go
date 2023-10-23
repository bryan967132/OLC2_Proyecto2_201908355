package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
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

func (w *While) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("---------- While ----------")

	continueLbl := c3dgen.NewLabel()
	c3dgen.AddLabel(continueLbl)

	condicion := w.Condition.Exec(Env, c3dgen)
	envWhile := env.NewEnv(Env, Env.Name+" While")

	envWhile.BreakLbl = condicion.FalseLabel
	envWhile.ContinueLbl = []string{continueLbl}

	for _, lbl := range condicion.TrueLabel {
		c3dgen.AddLabel(lbl)
	}

	w.Block.Exec(envWhile, c3dgen)
	c3dgen.AddGoto(continueLbl)

	for _, lbl := range condicion.FalseLabel {
		c3dgen.AddLabel(lbl)
	}

	c3dgen.AddComment("---------------------------")
	return nil
}
