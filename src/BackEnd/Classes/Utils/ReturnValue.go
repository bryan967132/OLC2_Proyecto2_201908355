package utils

type ReturnValue struct {
	StrValue   string
	IsTmp      bool
	Type       TypeExp
	TrueLabel  []interface{}
	FalseLabel []interface{}
	OutLabel   []interface{}
	IntValue   int
}

func NewValue(StrValue string, IsTmp bool, Type TypeExp) *ReturnValue {
	return &ReturnValue{StrValue: StrValue, IsTmp: IsTmp, Type: Type}
}
