package utils

type Arg struct {
	Line     int
	Column   int
	ExternID string
	Exp      interface{}
	IsInout  bool
}

func NewArg(line, column int, externID string, Exp interface{}, isInout bool) *Arg {
	return &Arg{line, column, externID, Exp, isInout}
}
