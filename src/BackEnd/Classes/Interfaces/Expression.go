package interfaces

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	utils "TSwift/Classes/Utils"
)

type Expression interface {
	LineN() int
	ColumnN() int
	Exec(env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue
}
