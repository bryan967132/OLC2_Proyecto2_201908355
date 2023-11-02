package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
	"fmt"
	"strconv"
	"strings"
)

type Primitive struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Value   interface{}
	Type    utils.Type
}

func NewPrimitive(line, column int, value interface{}, typeD utils.Type) *Primitive {
	return &Primitive{line, column, utils.PRIMITIVE, value, typeD}
}

func (pr *Primitive) LineN() int {
	return pr.Line
}

func (pr *Primitive) ColumnN() int {
	return pr.Column
}

func (pr *Primitive) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	switch pr.Type {
	case utils.INT:
		intValue, _ := strconv.Atoi(pr.Value.(string))
		return &utils.ReturnValue{StrValue: fmt.Sprintf("%v", pr.Value), NumValue: intValue, IsTmp: false, Type: pr.Type}
	case utils.FLOAT:
		fltValue, _ := strconv.ParseFloat(pr.Value.(string), 64)
		return &utils.ReturnValue{StrValue: fmt.Sprintf("%v", pr.Value), NumValue: fltValue, IsTmp: false, Type: pr.Type}
	case utils.BOOLEAN:
		trueLbl := c3dgen.NewLabel()
		falseLbl := c3dgen.NewLabel()
		if pr.Value.(string) == "true" {
			c3dgen.AddGoto(trueLbl)
		} else {
			c3dgen.AddGoto(falseLbl)
		}
		return &utils.ReturnValue{IsTmp: false, Type: pr.Type, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
	case utils.NIL:
		return &utils.ReturnValue{IsTmp: false, Type: pr.Type}
	default:
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\n", "\n")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\t", "\t")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\r", "\r")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\\"", "\"")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\'", "'")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\\\", "\\")
		newTemp := c3dgen.NewTemp()
		c3dgen.AddAssign(newTemp, "H")
		for _, asc := range []byte(pr.Value.(string)) {
			c3dgen.AddSetHeap("(int) H", strconv.Itoa(int(asc)))
			c3dgen.AddExpression("H", "H", "+", "1")
		}
		c3dgen.AddSetHeap("(int) H", "-1")
		c3dgen.AddExpression("H", "H", "+", "1")
		return &utils.ReturnValue{StrValue: newTemp, IsTmp: true, Type: pr.Type}
	}
}
