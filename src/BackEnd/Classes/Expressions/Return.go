package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
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

func (r *Return) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("--------- Return ----------")
	if r.Exp != nil {
		exp := r.Exp.Exec(env, c3dgen)
		if exp.Type == utils.BOOLEAN {
			newLabel := c3dgen.NewLabel()

			for _, lbl := range exp.TrueLabel {
				c3dgen.AddLabel(lbl)
			}
			c3dgen.AddSetStack("(int) P", "1")
			c3dgen.AddGoto(newLabel)

			for _, lbl := range exp.FalseLabel {
				c3dgen.AddLabel(lbl)
			}
			c3dgen.AddSetStack("(int) P", "0")

			c3dgen.AddLabel(newLabel)
		} else {
			c3dgen.AddSetStack("(int) P", exp.StrValue)
		}
		c3dgen.AddGoto(env.ReturnLbl)
		c3dgen.AddComment("---------------------------")
		return nil
	}
	c3dgen.AddSetStack("(int) P", "0")
	c3dgen.AddGoto(env.ReturnLbl)
	c3dgen.AddComment("---------------------------")
	return nil
}
