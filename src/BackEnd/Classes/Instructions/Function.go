package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
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

func (f *Function) GetType() utils.Type {
	return f.TypeRet
}

func (f *Function) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	variable := interface{}(f)
	Env.SaveFunction(f.ID, &variable, f.TypeRet, f.Line, f.Column)
	global := Env.GetGlobal()
	envFunc := env.NewEnv(global, "Funcion "+f.ID)

	envFunc.ReturnLbl = c3dgen.NewLabel()
	envFunc.Size = 1

	params := f.Params
	for i := 0; i < len(params); i++ {
		Type := params[i].Type.Value.(utils.Type)
		envFunc.SaveID(true, params[i].ID, &utils.ReturnValue{Type: Type}, Type, params[i].Line, params[i].Column)
	}

	c3dgen.AddFunc(f.ID)

	c3dgen.MainC3DCode = false
	f.Block.Exec(envFunc, c3dgen)

	c3dgen.AddGoto(envFunc.ReturnLbl)
	c3dgen.AddLabel(envFunc.ReturnLbl)
	c3dgen.MainC3DCode = true
	c3dgen.AddEnd()
	return nil
}
