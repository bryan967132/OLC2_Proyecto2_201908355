package env

import utils "TSwift/Classes/Utils"

type Symbol struct {
	IsVariable  bool
	IsPrimitive bool
	Value       interface{}
	Id          string
	Type        utils.Type
	ArrType     utils.Type
}

func NewSymbol(isVariable bool, isPrimitive bool, value interface{}, id string, Type, arrType utils.Type) *Symbol {
	return &Symbol{isVariable, isPrimitive, value, id, Type, arrType}
}
