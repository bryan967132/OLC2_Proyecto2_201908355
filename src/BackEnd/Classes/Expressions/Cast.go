package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cast struct {
	Line    int
	Column  int
	Destiny utils.Type
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewCast(line, column int, destiny utils.Type, exp interfaces.Expression) *Cast {
	return &Cast{line, column, destiny, utils.CAST, exp}
}

func (ct *Cast) LineN() int {
	return ct.Line
}

func (ct *Cast) ColumnN() int {
	return ct.Column
}

func (ct *Cast) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	value := ct.Exp.Exec(env, c3dgen)
	if ct.Destiny == utils.INT {
		if value.Type == utils.FLOAT {
			floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.NumValue), 64)
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpressionInit(newTemp, fmt.Sprintf("%v", floatValue))
			return &utils.ReturnValue{StrValue: newTemp, NumValue: int(floatValue), Type: ct.Destiny}
		}
		if value.Type == utils.STRING || value.Type == utils.CHAR {
			if ct.isValidNumericalString(fmt.Sprintf("%v", value.NumValue)) {
				floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.NumValue), 64)
				newTemp := c3dgen.NewTemp()
				c3dgen.AddExpressionInit(newTemp, fmt.Sprintf("%v", floatValue))
				return &utils.ReturnValue{StrValue: newTemp, NumValue: int(floatValue), Type: ct.Destiny}
			}
			env.SetError(fmt.Sprintf("La cadena \"%s\" no tiene formato numérico para castear a \"Int\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
			return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"Int\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	if ct.Destiny == utils.FLOAT {
		if value.Type == utils.STRING {
			if ct.isValidNumericalString(fmt.Sprintf("%v", value.NumValue)) {
				floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.NumValue), 64)
				newTemp := c3dgen.NewTemp()
				c3dgen.AddExpressionInit(newTemp, fmt.Sprintf("%v", floatValue))
				return &utils.ReturnValue{StrValue: newTemp, NumValue: floatValue, Type: ct.Destiny}
			}
			env.SetError(fmt.Sprintf("La cadena \"%s\" no tiene formato numérico para castear a \"Int\"", value.StrValue), ct.Exp.LineN(), ct.Exp.ColumnN())
			return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"Float\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	if ct.Destiny == utils.STRING {
		if value.Type == utils.INT {
			c3dgen.GenerateIntString()

			newTemp := c3dgen.NewTemp()

			c3dgen.AddExpression(newTemp, "P", "+", fmt.Sprintf("%v", env.Size))
			c3dgen.AddExpression(newTemp, newTemp, "+", "1")
			c3dgen.AddSetStack("(int) "+newTemp, value.StrValue)

			c3dgen.NewEnv(env.Size)
			c3dgen.AddCall("intString")

			newTemp2 := c3dgen.NewTemp()
			c3dgen.AddGetStack(newTemp2, "(int) P")
			c3dgen.PrevEnv(env.Size)

			return &utils.ReturnValue{StrValue: newTemp2, Type: ct.Destiny}
		}
		if value.Type == utils.FLOAT {
			newValue := (&Primitive{Value: fmt.Sprintf("%v", value.NumValue), Type: utils.STRING}).Exec(env, c3dgen)
			return &utils.ReturnValue{StrValue: newValue.StrValue, Type: ct.Destiny}
		}
		if value.Type == utils.BOOLEAN {
			newLabel := c3dgen.NewLabel()
			out := c3dgen.NewTemp()

			for _, lbl := range value.TrueLabel {
				c3dgen.AddLabel(lbl)
			}

			c3dgen.AddExpressionInit(out, "H")
			c3dgen.AddSetHeap("(int) H", "116")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "114")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "117")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "101")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "-1")
			c3dgen.AddExpression("H", "H", "+", "1")

			c3dgen.AddGoto(newLabel)

			for _, lbl := range value.FalseLabel {
				c3dgen.AddLabel(lbl)
			}

			c3dgen.AddExpressionInit(out, "H")
			c3dgen.AddSetHeap("(int) H", "102")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "97")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "108")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "115")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "101")
			c3dgen.AddExpression("H", "H", "+", "1")
			c3dgen.AddSetHeap("(int) H", "-1")
			c3dgen.AddExpression("H", "H", "+", "1")

			c3dgen.AddLabel(newLabel)

			return &utils.ReturnValue{StrValue: out, Type: ct.Destiny, IsTmp: true}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"String\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
	}
	env.SetError("Error: No hay casteo a \"Bool\" y \"Character\"", ct.Line, ct.Column)
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}

func (ct *Cast) isValidNumericalString(num string) bool {
	expresionRegular := `^\d+(\.\d+)?$`
	regexpValidador := regexp.MustCompile(expresionRegular)
	return regexpValidador.MatchString(num)
}

func (ct *Cast) cutDecimals(num string) string {
	parts := strings.Split(num, ".")
	if len(parts) == 2 {
		if len(parts[1]) > 5 {
			return parts[0] + "." + parts[1][:5]
		}
	} else if len(parts) == 1 {
		return num + ".0"
	}
	return num
}

func (ct *Cast) getType(Type utils.Type) string {
	switch Type {
	case utils.INT:
		return "Int"
	case utils.FLOAT:
		return "Float"
	case utils.STRING:
		return "String"
	case utils.BOOLEAN:
		return "Bool"
	case utils.CHAR:
		return "Character"
	default:
		return "nil"
	}
}
