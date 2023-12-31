package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"strconv"
)

type Print struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Exps     []interfaces.Expression
}

func NewPrint(line int, column int, exps []interfaces.Expression) *Print {
	return &Print{line, column, utils.PRINT, exps}
}

func (prt *Print) LineN() int {
	return prt.Line
}

func (prt *Print) ColumnN() int {
	return prt.Column
}

func (prt *Print) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("---------- Print ----------")
	if prt.Exps != nil {
		flag := false
		for _, exp := range prt.Exps {
			if flag {
				c3dgen.AddPrint(" ")
			}
			value1 := exp.Exec(env, c3dgen)
			if value1.Type == utils.INT {
				c3dgen.AddPrintf("d", fmt.Sprintf("(int) %s", value1.StrValue))
			} else if value1.Type == utils.FLOAT {
				c3dgen.AddPrintf("f", value1.StrValue)
			} else if value1.Type == utils.BOOLEAN {
				newLabel := c3dgen.NewLabel()
				for _, lbl := range value1.TrueLabel {
					c3dgen.AddLabel(lbl)
				}
				c3dgen.AddPrint("true")
				c3dgen.AddGoto(newLabel)
				for _, lbl := range value1.FalseLabel {
					c3dgen.AddLabel(lbl)
				}
				c3dgen.AddPrint("false")
				c3dgen.AddLabel(newLabel)
			} else if value1.Type == utils.STRING || value1.Type == utils.CHAR {
				c3dgen.GeneratePrintString()
				newTemp1 := c3dgen.NewTemp()
				newTemp2 := c3dgen.NewTemp()
				size := strconv.Itoa(env.Size)
				c3dgen.AddExpression(newTemp1, "P", "+", size)
				c3dgen.AddExpression(newTemp1, newTemp1, "+", "1")
				c3dgen.AddSetStack("(int) "+newTemp1, value1.StrValue)
				c3dgen.NewEnv(env.Size)
				c3dgen.AddCall("printString")
				c3dgen.AddGetStack(newTemp2, "(int) P")
				c3dgen.PrevEnv(env.Size)
			} else {
				c3dgen.AddPrint("nil")
			}
			flag = true
		}
	}
	c3dgen.AddPrint("\n")
	c3dgen.AddComment("---------------------------")
	return nil
}
