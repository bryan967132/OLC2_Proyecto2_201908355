package env

import utils "TSwift/Classes/Utils"

type Symbol struct {
	IsVariable  bool
	IsPrimitive bool
	Id          string
	Type        utils.Type
	ArrType     utils.Type
	Position    int
}

func NewSymbol(isVariable bool, isPrimitive bool, id string, Type, arrType utils.Type, position int) *Symbol {
	return &Symbol{isVariable, isPrimitive, id, Type, arrType, position}
}
