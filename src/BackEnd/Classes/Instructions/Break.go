package instructions

import (
	env "TSwift/Classes/Env"
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

func (b *Break) Exec(env *env.Env) *utils.ReturnType {
	return &utils.ReturnType{Value: b.TypeInst, Type: utils.NIL}
}
