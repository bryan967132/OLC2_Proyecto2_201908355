package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type CallFunction struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	ID      string
	Args    []utils.Arg
}

func NewCallFunction(line, column int, id string, args []utils.Arg) *CallFunction {
	return &CallFunction{line, column, utils.CALL_FUNC, id, args}
}

func (c *CallFunction) LineN() int {
	return c.Line
}

func (c *CallFunction) ColumnN() int {
	return c.Column
}

func (c *CallFunction) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	return nil
}
