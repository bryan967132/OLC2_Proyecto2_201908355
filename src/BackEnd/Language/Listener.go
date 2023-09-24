package language

import (
	parser "TSwift/Language/Parser"
)

type TSwfitListener struct {
	*parser.BaseParserListener
	Code []interface{}
}

func NewTSwfitListener() *TSwfitListener {
	return new(TSwfitListener)
}

func (this *TSwfitListener) ExitInit(ctx *parser.InitContext) {
	this.Code = ctx.GetResult()
}
