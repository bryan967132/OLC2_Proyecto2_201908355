package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Relational struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewRelational(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Relational {
	return &Relational{line, column, utils.RELATIONAL_OP, exp1, sign, exp2}
}

func (rl *Relational) LineN() int {
	return rl.Line
}

func (rl *Relational) ColumnN() int {
	return rl.Column
}

func (rl *Relational) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	switch rl.Sign {
	case "==":
		return rl.equal(env, c3dgen)
	case "!=":
		return rl.notEqual(env, c3dgen)
	case ">=":
		return rl.moreEqual(env, c3dgen)
	case "<=":
		return rl.lessEqual(env, c3dgen)
	case ">":
		return rl.more(env, c3dgen)
	case "<":
		return rl.less(env, c3dgen)
	default:
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
}

func (rl *Relational) equal(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.BOOLEAN && value2.Type == utils.BOOLEAN ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, "==", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (rl *Relational) notEqual(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.BOOLEAN && value2.Type == utils.BOOLEAN ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, "!=", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (rl *Relational) moreEqual(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, ">=", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (rl *Relational) lessEqual(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, "<=", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (rl *Relational) more(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, ">", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (rl *Relational) less(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := rl.Exp1.Exec(env, c3dgen)
	value2 := rl.Exp2.Exec(env, c3dgen)
	if value1.Type == utils.INT && value2.Type == utils.INT ||
		value1.Type == utils.FLOAT && value2.Type == utils.FLOAT ||
		value1.Type == utils.STRING && value2.Type == utils.STRING ||
		value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		c3dgen.AddIf(value1.StrValue, value2.StrValue, "<", trueLbl)
		c3dgen.AddGoto(falseLbl)
		return &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	}
	env.SetError("Los tipos no son válidos para operaciones relacionales", rl.Exp1.LineN(), rl.Exp1.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}
