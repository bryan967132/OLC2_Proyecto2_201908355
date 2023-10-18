package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type Break struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
}

func NewBreak(line, column int) *Break {
	return &Break{line, column, utils.BREAK}
}

func (b *Break) LineN() int {
	return b.Line
}

func (b *Break) ColumnN() int {
	return b.Column
}

func (b *Break) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
