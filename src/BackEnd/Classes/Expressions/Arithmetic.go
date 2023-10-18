package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"strconv"
)

type Arithmetic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
	Type    utils.Type
}

func NewArithmetic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Arithmetic {
	return &Arithmetic{line, column, utils.ARITHMETIC_OP, exp1, sign, exp2, utils.NIL}
}

func (ar *Arithmetic) LineN() int {
	return ar.Line
}

func (ar *Arithmetic) ColumnN() int {
	return ar.Column
}

func (ar *Arithmetic) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	switch ar.Sign {
	case "+":
		return ar.plus(env, c3dgen)
	case "-":
		if ar.Exp1 != nil {
			return ar.minus(env, c3dgen)
		}
		return ar.uminus(env, c3dgen)
	case "*":
		return ar.mult(env, c3dgen)
	case "/":
		return ar.div(env, c3dgen)
	case "%":
		return ar.mod(env, c3dgen)
	default:
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
}

func (ar *Arithmetic) plus(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := ar.Exp1.Exec(env, c3dgen)
	if value1.Type == utils.BOOLEAN {
		env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp1.LineN(), ar.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	if value1.Type == utils.CHAR {
		value1.Type = utils.STRING
	}
	value2 := ar.Exp2.Exec(env, c3dgen)
	if value2.Type == utils.CHAR {
		value2.Type = utils.STRING
	}
	if int(value1.Type) < len(utils.Plus) && int(value2.Type) < len(utils.Plus[0]) {
		ar.Type = utils.Plus[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "+", value2.StrValue)
			intValue1, _ := strconv.Atoi(fmt.Sprintf("%v", value1.NumValue))
			intValue2, _ := strconv.Atoi(fmt.Sprintf("%v", value2.NumValue))
			return &utils.ReturnValue{StrValue: newTemp, NumValue: intValue1 + intValue2, IsTmp: true, Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "+", value2.StrValue)
			fltValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.StrValue), 64)
			fltValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.StrValue), 64)
			return &utils.ReturnValue{StrValue: newTemp, NumValue: fltValue1 + fltValue2, IsTmp: true, Type: ar.Type}
		}
		if ar.Type == utils.STRING {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddAssign(newTemp, "H")
			for _, asc := range []byte(fmt.Sprintf("%v%v", value1.StrValue, value2.StrValue)) {
				c3dgen.AddSetHeap("(int) H", strconv.Itoa(int(asc)))
				c3dgen.AddExpression("H", "H", "+", "1")
			}
			c3dgen.AddSetHeap("(int) H", "-1")
			c3dgen.AddExpression("H", "H", "+", "1")
			return &utils.ReturnValue{StrValue: fmt.Sprintf("%v%v", value1.StrValue, value2.StrValue), IsTmp: true, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ar *Arithmetic) minus(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := ar.Exp1.Exec(env, c3dgen)
	if value1.Type == utils.STRING || value1.Type == utils.BOOLEAN || value1.Type == utils.CHAR {
		env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp1.LineN(), ar.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	value2 := ar.Exp2.Exec(env, c3dgen)
	if int(value1.Type) < len(utils.Minus) && int(value2.Type) < len(utils.Minus[0]) {
		ar.Type = utils.Minus[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "-", value2.StrValue)
			intValue1, _ := strconv.Atoi(fmt.Sprintf("%v", value1.NumValue))
			intValue2, _ := strconv.Atoi(fmt.Sprintf("%v", value2.NumValue))
			return &utils.ReturnValue{StrValue: newTemp, NumValue: intValue1 - intValue2, IsTmp: true, Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "-", value2.StrValue)
			fltValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.StrValue), 64)
			fltValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.StrValue), 64)
			return &utils.ReturnValue{StrValue: newTemp, NumValue: fltValue1 - fltValue2, IsTmp: true, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ar *Arithmetic) uminus(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value2 := ar.Exp2.Exec(env, c3dgen)
	ar.Type = value2.Type
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, "", "+", value2.StrValue)
			intValue2, _ := strconv.Atoi(fmt.Sprintf("%v", value2.NumValue))
			return &utils.ReturnValue{StrValue: newTemp, NumValue: -intValue2, IsTmp: true, Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, "", "-", value2.StrValue)
			floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.StrValue), 64)
			return &utils.ReturnValue{StrValue: newTemp, NumValue: -floatValue2, IsTmp: true, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ar *Arithmetic) mult(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := ar.Exp1.Exec(env, c3dgen)
	if value1.Type == utils.STRING || value1.Type == utils.BOOLEAN || value1.Type == utils.CHAR {
		env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp1.LineN(), ar.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	value2 := ar.Exp2.Exec(env, c3dgen)
	if int(value1.Type) < len(utils.Mult) && int(value2.Type) < len(utils.Mult[0]) {
		ar.Type = utils.Mult[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "*", value2.StrValue)
			intValue1, _ := strconv.Atoi(fmt.Sprintf("%v", value1.NumValue))
			intValue2, _ := strconv.Atoi(fmt.Sprintf("%v", value2.NumValue))
			return &utils.ReturnValue{StrValue: newTemp, NumValue: intValue1 * intValue2, IsTmp: true, Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, value1.StrValue, "*", value2.StrValue)
			fltValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.StrValue), 64)
			fltValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.StrValue), 64)
			return &utils.ReturnValue{StrValue: newTemp, NumValue: fltValue1 * fltValue2, IsTmp: true, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ar *Arithmetic) div(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := ar.Exp1.Exec(env, c3dgen)
	if value1.Type == utils.STRING || value1.Type == utils.BOOLEAN || value1.Type == utils.CHAR {
		env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp1.LineN(), ar.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	value2 := ar.Exp2.Exec(env, c3dgen)
	if int(value1.Type) < len(utils.Div) && int(value2.Type) < len(utils.Div[0]) {
		ar.Type = utils.Div[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT || ar.Type == utils.FLOAT {
			newLbl1 := c3dgen.NewLabel()
			newLbl2 := c3dgen.NewLabel()
			newTemp1 := c3dgen.NewTemp()
			c3dgen.AddExpressionInit(newTemp1, value2.StrValue)
			c3dgen.AddIf(newTemp1, "0", "!=", newLbl1)
			c3dgen.AddPrint("MathError")
			newTemp2 := c3dgen.NewTemp()
			c3dgen.AddExpressionInit(newTemp2, "0")
			c3dgen.AddGoto(newLbl2)
			c3dgen.AddLabel(newLbl1)
			c3dgen.AddExpression(newTemp2, value1.StrValue, "/", newTemp1)
			c3dgen.AddLabel(newLbl2)
			if value2.NumValue.(int) == 0 {
				env.SetError("División entre cero", ar.Exp2.LineN(), ar.Exp2.ColumnN())
			}
			return &utils.ReturnValue{StrValue: newTemp2, IsTmp: true, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ar *Arithmetic) mod(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value1 := ar.Exp1.Exec(env, c3dgen)
	if value1.Type == utils.STRING || value1.Type == utils.BOOLEAN || value1.Type == utils.CHAR {
		env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp1.LineN(), ar.Exp1.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	value2 := ar.Exp2.Exec(env, c3dgen)
	if int(value1.Type) < len(utils.Div) && int(value2.Type) < len(utils.Div[0]) {
		ar.Type = utils.Div[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT || ar.Type == utils.FLOAT {
			newLbl1 := c3dgen.NewLabel()
			newLbl2 := c3dgen.NewLabel()
			c3dgen.AddIf(value2.StrValue, "0", "!=", newLbl1)
			c3dgen.AddPrint("MathError")
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpressionInit(newTemp, "0")
			c3dgen.AddGoto(newLbl2)
			c3dgen.AddLabel(newLbl1)
			c3dgen.AddExpression(newTemp, value1.StrValue, "%", value2.StrValue)
			c3dgen.AddLabel(newLbl2)
			if value2.NumValue.(int) == 0 {
				env.SetError("División entre cero", ar.Exp2.LineN(), ar.Exp2.ColumnN())
			}
			return &utils.ReturnValue{StrValue: newTemp, Type: ar.Type}
		}
	}
	env.SetError("Los tipos no son válidos para operaciones aritméticas", ar.Exp2.LineN(), ar.Exp2.ColumnN())
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}
