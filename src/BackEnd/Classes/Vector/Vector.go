package vector

import (
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Vector struct {
	IsMatrix    bool
	IsReuseID   bool
	IsRepeating bool
	Length      int
	Dims        int
	Type        *utils.AttribsType
	Vectors     []*Vector
	Elements    []interfaces.Expression
	Values      []*utils.ReturnValue
	ID          string
	Repeating   *Repeating
}

func NewMatrix(Type *utils.AttribsType, Vectors []*Vector) *Vector {
	return &Vector{IsMatrix: true, IsReuseID: false, IsRepeating: false, Type: Type, Vectors: Vectors}
}

func NewVector(Type *utils.AttribsType, Elements []interfaces.Expression) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: false, IsRepeating: false, Type: Type, Elements: Elements}
}

func NewMatrixRepeating(repeating *Repeating) *Vector {
	return &Vector{IsMatrix: true, IsRepeating: true, Repeating: repeating}
}

func NewReuseVector(id string) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: true, IsRepeating: false, ID: id}
}
