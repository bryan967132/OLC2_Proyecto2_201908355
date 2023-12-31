package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Case struct {
	Line         int
	Column       int
	TypeInst     utils.TypeInst
	Case         interfaces.Expression
	Block        interfaces.Instruction
	CaseEvaluate *utils.ReturnValue
	Flag         bool
}

func NewCase(line, column int, case_ interfaces.Expression, block interfaces.Instruction) *Case {
	return &Case{Line: line, Column: column, TypeInst: utils.CASE, Case: case_, Block: block, Flag: false}
}

func (c *Case) LineN() int {
	return c.Line
}

func (c *Case) ColumnN() int {
	return c.Column
}

func (c *Case) SetCase(caseEvaluate *utils.ReturnValue) {
	c.CaseEvaluate = caseEvaluate
}

func (c *Case) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
