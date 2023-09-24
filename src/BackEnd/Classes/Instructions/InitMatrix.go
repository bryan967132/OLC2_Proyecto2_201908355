package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type InitMatrix struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       *utils.VectorType
	Value      *vector.Vector
}

func NewInitMatrix(line, column int, isVar bool, id string, Type *utils.VectorType, value *vector.Vector) *InitMatrix {
	return &InitMatrix{line, column, isVar, utils.INIT_ID, id, Type, value}
}

func (in *InitMatrix) LineN() int {
	return in.Line
}

func (in *InitMatrix) ColumnN() int {
	return in.Column
}

func (in *InitMatrix) Exec(env *env.Env) *utils.ReturnType {
	return nil
}

func (in *InitMatrix) getType(Type utils.Type) string {
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
