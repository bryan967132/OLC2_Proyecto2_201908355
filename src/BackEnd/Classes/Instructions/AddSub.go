package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type AddSub struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Sign     string
	Exp      interfaces.Expression
}

func NewAddSub(line, column int, id, sign string, exp interfaces.Expression) *AddSub {
	if id == "+=" {
		return &AddSub{line, column, utils.ADD, id, sign, exp}
	}
	return &AddSub{line, column, utils.SUB, id, sign, exp}
}

func (a *AddSub) LineN() int {
	return a.Line
}

func (a *AddSub) ColumnN() int {
	return a.Column
}

func (a *AddSub) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
