package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type Continue struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
}

func NewContinue(line, column int) *Continue {
	return &Continue{line, column, utils.CONTINUE}
}

func (c *Continue) LineN() int {
	return c.Line
}

func (c *Continue) ColumnN() int {
	return c.Column
}

func (c *Continue) Exec(env *env.Env) *utils.ReturnType {
	return &utils.ReturnType{Value: c.TypeInst, Type: utils.NIL}
}
