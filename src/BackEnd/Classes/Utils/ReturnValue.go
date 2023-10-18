package utils

type ReturnValue struct {
	StrValue   string
	IntValue   int
	IsTmp      bool
	Type       Type
	TrueLabel  []string
	FalseLabel []string
	OutLabel   []string
}
