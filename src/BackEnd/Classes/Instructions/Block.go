package instructions

import (
	env "TSwift/Classes/Env"
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

func (bk *Block) Exec(Env *env.Env) *utils.ReturnValue {
	return nil
}
