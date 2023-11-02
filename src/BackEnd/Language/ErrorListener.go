package language

import (
	"github.com/antlr4-go/antlr/v4"
)

type CustomSyntaxError struct {
	Line   int
	Column int
	Msg    string
}

type TSwfitErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*CustomSyntaxError
}

func NewTSwfitErrorListener() *TSwfitErrorListener {
	return new(TSwfitErrorListener)
}

func (c *TSwfitErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		Line:   line,
		Column: column,
		Msg:    msg,
	})
}
