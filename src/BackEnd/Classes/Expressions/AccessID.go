package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
	"strconv"
)

type AccessID struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Id      string
}

func NewAccessID(line, column int, id string) *AccessID {
	return &AccessID{line, column, utils.ACCESS_ID, id}
}

func (ac *AccessID) LineN() int {
	return ac.Line
}

func (ac *AccessID) ColumnN() int {
	return ac.Column
}

func (ac *AccessID) Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	c3dgen.AddComment("--------- Acceso ----------")
	value := env.GetValueID(ac.Id, ac.Line, ac.Column)
	if value != nil {
		newTemp := c3dgen.NewTemp()
		c3dgen.AddGetStack(newTemp, strconv.Itoa(value.Position))
		if value.Type == utils.BOOLEAN {
			trueLbl := c3dgen.NewLabel()
			falseLbl := c3dgen.NewLabel()
			c3dgen.AddIf(newTemp, "1", "==", trueLbl)
			c3dgen.AddGoto(falseLbl)
			result := &utils.ReturnValue{IsTmp: false, Type: utils.BOOLEAN, TrueLabel: []string{trueLbl}, FalseLabel: []string{falseLbl}}
			c3dgen.AddComment("---------------------------")
			return result
		}
		result := &utils.ReturnValue{StrValue: newTemp, IsTmp: true, Type: value.Type}
		c3dgen.AddComment("---------------------------")
		return result
	}
	c3dgen.AddComment("---------------------------")
	return &utils.ReturnValue{IsTmp: false, Type: utils.NIL}
}
