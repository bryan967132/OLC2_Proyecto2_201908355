package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Function struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
	Params   []utils.Param
	Block    interfaces.Instruction
	TypeRet  utils.Type
}

func NewFunction(line, column int, id string, params []utils.Param, Block interfaces.Instruction, TypeRet utils.Type) *Function {
	return &Function{line, column, utils.INIT_FUNCTION, id, params, Block, TypeRet}
}

func (f *Function) LineN() int {
	return f.Line
}

func (f *Function) ColumnN() int {
	return f.Column
}

func (f *Function) GetParams() []utils.Param {
	return f.Params
}

func (f *Function) GetBlock() interfaces.Instruction {
	return f.Block
}

func (f *Function) Exec(Env *env.Env) *utils.ReturnValue {
	return nil
}
