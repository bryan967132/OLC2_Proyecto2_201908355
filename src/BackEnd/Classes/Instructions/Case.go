package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Case struct {
	Line         int
	Column       int
	TypeInst     utils.TypeInst
	Case         interfaces.Expression
	Block        interfaces.Instruction
	CaseEvaluate *utils.ReturnType
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

func (c *Case) SetCase(caseEvaluate *utils.ReturnType) {
	c.CaseEvaluate = caseEvaluate
}

func (c *Case) Exec(Env *env.Env) *utils.ReturnType {
	return nil
}
