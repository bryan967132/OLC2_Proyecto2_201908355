package utils

import "fmt"

type ReturnType struct {
	Value interface{}
	Type  Type
}

func (ret *ReturnType) ToString() string {
	return fmt.Sprintf("Object: %v, Type: %v", ret.Value, ret.Type)
}

func (ret *ReturnType) GetCopy() *ReturnType {
	return &ReturnType{Value: ret.Value, Type: ret.Type}
}
