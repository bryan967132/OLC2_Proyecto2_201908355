package vector

import (
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Repeating struct {
	Length    int
	Dims      int
	Type      utils.Type
	Value     interfaces.Expression
	Times     interfaces.Expression
	Repeating *Repeating
}

func NewRepeating(length, dims int, Type utils.Type, value, times interfaces.Expression, repeating *Repeating) *Repeating {
	return &Repeating{length, dims, Type, value, times, repeating}
}
