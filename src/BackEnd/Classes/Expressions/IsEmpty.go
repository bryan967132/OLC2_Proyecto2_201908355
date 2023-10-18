package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type IsEmpty struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewIsEmpty(line, column int, exp interfaces.Expression) *IsEmpty {
	return &IsEmpty{line, column, utils.ARRAY_ISEMPTY, exp}
}

func (i *IsEmpty) LineN() int {
	return i.Line
}

func (i *IsEmpty) ColumnN() int {
	return i.Column
}

func (i *IsEmpty) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
