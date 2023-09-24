package interfaces

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type Function interface {
	LineN() int
	ColumnN() int
	GetParams() []utils.Param
	GetBlock() Instruction
	Exec(Env *env.Env) *utils.ReturnType
}
