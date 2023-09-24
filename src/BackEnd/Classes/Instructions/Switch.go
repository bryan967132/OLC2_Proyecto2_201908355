package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Switch struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Arg      interfaces.Expression
	Cases    interface{}
	Default  interfaces.Instruction
}

func NewSwitch(line, column int, arg interfaces.Expression, cases interface{}, _default interfaces.Instruction) *Switch {
	return &Switch{line, column, utils.SWITCH, arg, cases, _default}
}

func (s *Switch) LineN() int {
	return s.Line
}

func (s *Switch) ColumnN() int {
	return s.Column
}

func (s *Switch) Exec(Env *env.Env) *utils.ReturnType {
	return nil
}
