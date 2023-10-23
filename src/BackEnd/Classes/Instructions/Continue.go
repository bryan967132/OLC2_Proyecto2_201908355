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
	Tag      string
}

func NewContinue(line, column int) *Continue {
	return &Continue{Line: line, Column: column, TypeInst: utils.CONTINUE}
}

func (c *Continue) LineN() int {
	return c.Line
}

func (c *Continue) ColumnN() int {
	return c.Column
}

func (c *Continue) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("-------- Continue ---------")
	for _, lbl := range env.ContinueLbl {
		c3dgen.AddGoto(lbl)
	}
	c3dgen.AddComment("---------------------------")
	return nil
}
