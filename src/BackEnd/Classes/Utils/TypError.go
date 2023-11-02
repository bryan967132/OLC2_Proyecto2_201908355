package utils

type TypeError string

const (
	LEXICAL  TypeError = "LEXICO"
	SYNTAX   TypeError = "SINTACTICO"
	SEMANTIC TypeError = "SEMANTICO"
)
