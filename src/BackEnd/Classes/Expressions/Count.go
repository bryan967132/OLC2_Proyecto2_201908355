package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Count struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewCount(line, column int, exp interfaces.Expression) *Count {
	return &Count{line, column, utils.ARRAY_COUNT, exp}
}

func (c *Count) LineN() int {
	return c.Line
}

func (c *Count) ColumnN() int {
	return c.Column
}

func (c *Count) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
