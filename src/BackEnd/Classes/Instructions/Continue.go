package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
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

func (c *Continue) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
