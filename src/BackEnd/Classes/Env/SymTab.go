package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type SymTab struct {
	Num         int
	Line        int
	Column      int
	IsVariable  bool
	IsPrimitive bool
	Id          string
	NameEnv     string
	Type        utils.Type
	ArrType     utils.Type
}

func NewSymTab(line, column int, isVariable bool, isPrimitive bool, id, nameEnv string, Type, arrType utils.Type) SymTab {
	return SymTab{0, line, column, isVariable, isPrimitive, id, nameEnv, Type, arrType}
}

func (s *SymTab) ToString() string {
	return fmt.Sprintf("Identificador: %s, Tipo: %v %v, Entorno: %s. %v:%v", s.Id, s.getType(s.Type), s.getType(s.ArrType), s.NameEnv, s.Line, s.Column)
}

func (s *SymTab) Hash() string {
	return fmt.Sprintf("%v_%v_%v_%v_%v_%v_%v_%v", s.Id, s.Type, s.ArrType, s.NameEnv, s.Line, s.Column, s.IsVariable, s.IsPrimitive)
}

func (s *SymTab) GetDot() string {
	if s.IsPrimitive || s.IsVariable {
		if s.IsPrimitive {
			if s.IsVariable {
				return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Variable</td><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.NameEnv, s.Line, s.Column)
			}
			return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Constante</td><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.NameEnv, s.Line, s.Column)
		}
		if s.IsVariable {
			return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Variable</td><td bgcolor="white">%v %v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.getType(s.ArrType), s.NameEnv, s.Line, s.Column)
		}
		return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Constante</td><td bgcolor="white">%v %v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.getType(s.ArrType), s.NameEnv, s.Line, s.Column)
	}
	if s.Type != utils.NIL {
		return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Función</td><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.NameEnv, s.Line, s.Column)
	}
	return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">Método</td><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, s.Num, s.Id, s.getType(s.Type), s.NameEnv, s.Line, s.Column)
}

func (s *SymTab) getType(Type utils.Type) string {
	switch Type {
	case utils.INT:
		return "Int"
	case utils.FLOAT:
		return "Float"
	case utils.STRING:
		return "String"
	case utils.BOOLEAN:
		return "Bool"
	case utils.CHAR:
		return "Character"
	case utils.VECTOR:
		return "Vector"
	case utils.MATRIX:
		return "Matriz"
	default:
		return "nil"
	}
}
