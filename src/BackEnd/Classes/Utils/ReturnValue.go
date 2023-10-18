package utils

type ReturnValue struct {
	StrValue   string
	NumValue   interface{}
	IsTmp      bool
	Type       Type
	TrueLabel  []string
	FalseLabel []string
	OutLabel   []string
}
