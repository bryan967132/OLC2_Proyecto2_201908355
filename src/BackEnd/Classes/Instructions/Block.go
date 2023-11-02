package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Block struct {
	Line         int
	Column       int
	TypeInst     utils.TypeInst
	Instructions []interface{}
}

func NewBlock(line, column int, instructions []interface{}) *Block {
	return &Block{line, column, utils.BLOCK_INST, instructions}
}

func (bk *Block) LineN() int {
	return bk.Line
}

func (bk *Block) ColumnN() int {
	return bk.Column
}

func (bk *Block) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	newEnv := env.NewEnv(Env, Env.Name)
	newEnv.ContinueLbl = Env.ContinueLbl
	newEnv.BreakLbl = Env.BreakLbl
	newEnv.ReturnLbl = Env.ReturnLbl
	newEnv.Size = Env.Size
	var ret *utils.ReturnValue
	var inst interfaces.Instruction
	outlbl := []string{}
	for _, instruction := range bk.Instructions {
		inst = instruction.(interfaces.Instruction)
		ret = inst.Exec(newEnv, c3dgen)
		if ret != nil {
			for _, lbl := range ret.OutLabel {
				outlbl = append(outlbl, lbl)
			}
			return ret
		}
	}
	return &utils.ReturnValue{OutLabel: outlbl}
}
