package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Logic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewLogic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Logic {
	return &Logic{line, column, utils.LOGIC_OP, exp1, sign, exp2}
}

func (lg *Logic) LineN() int {
	return lg.Line
}

func (lg *Logic) ColumnN() int {
	return lg.Column
}

func (lg *Logic) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	switch lg.Sign {
	case "&&":
		return lg.and(env, c3dgen)
	case "||":
		return lg.or(env, c3dgen)
	case "!":
		return lg.not(env, c3dgen)
	default:
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
}

func (lg *Logic) and(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := lg.Exp1.Exec(env, c3dgen)
	if value1.Type != utils.BOOLEAN {
		env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp1.LineN(), lg.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	for _, lbl := range value1.TrueLabel {
		c3dgen.AddLabel(lbl)
	}
	value2 := lg.Exp2.Exec(env, c3dgen)
	if value2.Type == utils.BOOLEAN {
		result := &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN}
		result.TrueLabel = append(value2.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(value2.FalseLabel, result.FalseLabel...)
		result.FalseLabel = append(value1.FalseLabel, result.FalseLabel...)
		return result
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (lg *Logic) or(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := lg.Exp1.Exec(env, c3dgen)
	if value1.Type != utils.BOOLEAN {
		env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp1.LineN(), lg.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	for _, lbl := range value1.FalseLabel {
		c3dgen.AddLabel(lbl)
	}
	value2 := lg.Exp2.Exec(env, c3dgen)
	if value2.Type == utils.BOOLEAN {
		result := &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN}
		result.TrueLabel = append(value1.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(value2.FalseLabel, result.FalseLabel...)
		result.TrueLabel = append(value2.TrueLabel, result.TrueLabel...)
		return result
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (lg *Logic) not(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value2 := lg.Exp2.Exec(env, c3dgen)
	if value2.Type == utils.BOOLEAN {
		result := &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN}
		result.TrueLabel = append(value2.FalseLabel, result.TrueLabel...)
		result.FalseLabel = append(value2.TrueLabel, result.FalseLabel...)
		return result
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}
