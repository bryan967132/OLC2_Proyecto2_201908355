package interfaces

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type Function interface {
	LineN() int
	ColumnN() int
	GetParams() []utils.Param
	GetBlock() Instruction
	Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue
}
