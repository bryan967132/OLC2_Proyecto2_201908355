package interfaces

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type Expression interface {
	LineN() int
	ColumnN() int
	Exec(env *env.Env) *utils.ReturnType
}
