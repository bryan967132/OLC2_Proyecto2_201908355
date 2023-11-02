package utils

type TypeExp string

const (
	PRIMITIVE     TypeExp = "PRIMITIVE"
	ARITHMETIC_OP TypeExp = "ARITHMETIC_OP"
	LOGIC_OP      TypeExp = "LOGIC_OP"
	RELATIONAL_OP TypeExp = "RELATIONAL_OP"
	ACCESS_ID     TypeExp = "ACCESS_ID"
	INC           TypeExp = "INC"
	DEC           TypeExp = "DEC"
	PARAMETER     TypeExp = "PARAMETER"
	CALL_FUNC     TypeExp = "CALL_FUNC"
	RETURN        TypeExp = "RETURN"
	CAST          TypeExp = "CAST"
	ACCESS_ARRAY  TypeExp = "ACCESS_ARRAY"
	ARRAY_COUNT   TypeExp = "ACCESS_COUNT"
	ARRAY_ISEMPTY TypeExp = "ACCESS_ISEMPTY"
)
