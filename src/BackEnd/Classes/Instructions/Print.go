package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Print struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Exps     []interfaces.Expression
}

func NewPrint(line int, column int, exps []interfaces.Expression) *Print {
	return &Print{line, column, utils.PRINT, exps}
}

func (prt *Print) LineN() int {
	return prt.Line
}

func (prt *Print) ColumnN() int {
	return prt.Column
}

func (prt *Print) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}