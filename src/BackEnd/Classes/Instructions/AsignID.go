package instructions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"strconv"
)

type AsignID struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Value    interfaces.Expression
}

func NewAsignID(line, column int, id string, value interfaces.Expression) *AsignID {
	return &AsignID{line, column, utils.ASIGN_ID, id, value}
}

func (as *AsignID) LineN() int {
	return as.Line
}

func (as *AsignID) ColumnN() int {
	return as.Column
}

func (as *AsignID) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("------- Asignacion --------")
	variable := env.GetValueID(as.Id, as.Line, as.Column)
	if variable != nil {
		value := as.Value.Exec(env, c3dgen)
		c3dgen.AddSetStack(strconv.Itoa(variable.Position), value.StrValue)
		c3dgen.AddComment("---------------------------")
		return nil
	}
	c3dgen.AddComment("---------------------------")
	env.SetError("Resignación de valor a variable inexistente", as.Line, as.Column)
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}
