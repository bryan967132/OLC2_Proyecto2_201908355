package utils

type AttribsType struct {
	Line        int
	Column      int
	Value       interface{}
	IsPrimitive bool
}

func NewAttribsType(line, column int, value interface{}, isPrimitive bool) *AttribsType {
	return &AttribsType{line, column, value, isPrimitive}
}
