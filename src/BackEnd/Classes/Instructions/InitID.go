package instructions

import (
	env "TSwift/Classes/Env"
	expressions "TSwift/Classes/Expressions"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type InitID struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       utils.Type
	Value      interfaces.Expression
}

func NewInitID(line, column int, isVariable bool, id string, Type utils.Type, value interfaces.Expression) *InitID {
	return &InitID{line, column, isVariable, utils.INIT_ID, id, Type, value}
}

func (in *InitID) LineN() int {
	return in.Line
}

func (in *InitID) ColumnN() int {
	return in.Column
}

func (in *InitID) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	if in.Value != nil {
		value := in.Value.Exec(env, c3dgen)
		if in.Type != utils.NIL {
			if in.Type == value.Type || in.Type == utils.STRING && value.Type == utils.CHAR || in.Type == utils.FLOAT && value.Type == utils.INT {
				in.Init(value, env, c3dgen)
				return nil
			}
			env.SetError("Los tipos no coinciden en la declaración", in.Line, in.Column)
			return nil
		}
		in.Init(value, env, c3dgen)
	} else {
		in.Init(nil, env, c3dgen)
	}
	return nil
}

func (in *InitID) Init(value *utils.ReturnValue, Env *env.Env, c3dgen *C3DGen.C3DGen) {
	c3dgen.AddComment("------- Declaración -------")
	var newID *env.Symbol
	if value != nil {
		newID = Env.SaveID(in.IsVariable, in.Id, &utils.ReturnValue{Type: value.Type}, value.Type, in.Line, in.Column)
		if value.Type == utils.BOOLEAN {
			newLabel := c3dgen.NewLabel()
			for _, lbl := range value.TrueLabel {
				c3dgen.AddLabel(lbl)
			}
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, "P", "+", fmt.Sprintf("%v", newID.Position))
			c3dgen.AddSetStack(fmt.Sprintf("(int) %v", newTemp), "1")
			c3dgen.AddGoto(newLabel)
			for _, lbl := range value.FalseLabel {
				c3dgen.AddLabel(lbl)
			}
			newTemp = c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, "P", "+", fmt.Sprintf("%v", newID.Position))
			c3dgen.AddSetStack(fmt.Sprintf("(int) %v", newTemp), "0")
			c3dgen.AddGoto(newLabel)
			c3dgen.AddLabel(newLabel)
		} else {
			newTemp := c3dgen.NewTemp()
			c3dgen.AddExpression(newTemp, "P", "+", fmt.Sprintf("%v", newID.Position))
			c3dgen.AddSetStack(fmt.Sprintf("(int) %v", newTemp), value.StrValue)
		}
	} else {
		newID = Env.SaveID(in.IsVariable, in.Id, &utils.ReturnValue{Type: utils.NIL}, in.Type, in.Line, in.Column)
		nilValue := *expressions.NewPrimitive(0, 0, "nil", utils.STRING).Exec(Env, c3dgen)
		newTemp := c3dgen.NewTemp()
		c3dgen.AddExpression(newTemp, "P", "+", fmt.Sprintf("%v", newID.Position))
		c3dgen.AddSetStack(fmt.Sprintf("(int) %v", newTemp), nilValue.StrValue)
	}
	c3dgen.AddComment("---------------------------")
}
