package controller

import (
	env "TSwift/Classes/Env"
	instructions "TSwift/Classes/Instructions"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

type Parser struct {
	Code string `json:"code"`
}

func (c Controller) Running(ctx *fiber.Ctx) error {
	return ctx.SendString("Interpreter is running!!!")
}

var ParseTree antlr.Tree
var ParserG *parser.ParserParser

func (c Controller) Parser(ctx *fiber.Ctx) error {
	var reqBody Parser
	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.JSON(fiber.Map{
			"console": "¡Ha ocurrido un error!",
		})
	}

	inputStream := antlr.NewInputStream(reqBody.Code)
	scanner := parser.NewScanner(inputStream)
	scannerErrors := listener.NewTSwfitErrorListener()
	scanner.RemoveErrorListeners()
	scanner.AddErrorListener(scannerErrors)

	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(tokens)

	parserErrors := listener.NewTSwfitErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.TSwfitListener = listener.NewTSwfitListener()

	global := env.NewEnv(nil, "Global")
	func() {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(Replaces(fmt.Sprintf("%v", r)), 0, 0)
			}
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
			ParseTree = tree
			ParserG = parser

			env.SymTable = env.NewSymbolTable()
			utils.PrintConsole = []string{}
			utils.Errors = []utils.Error{}

			for _, fail := range scannerErrors.Errors {
				utils.Errors = append(utils.Errors, *utils.NewError(fail.Line, fail.Column, utils.LEXICAL, Replaces(fail.Msg)))
			}

			for _, fail := range parserErrors.Errors {
				utils.Errors = append(utils.Errors, *utils.NewError(fail.Line, fail.Column, utils.SYNTAX, Replaces(fail.Msg)))
			}

			execute := listener.Code
			for _, instruction := range execute {
				func() {
					defer func() {
						if r := recover(); r != nil {
							global.SetError(Replaces(fmt.Sprintf("%v", r)), instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN())
						}
					}()
					if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); ok {
						instruction.(interfaces.Instruction).Exec(global)
					}
				}()
			}
			for _, instruction := range execute {
				func() {
					defer func() {
						if r := recover(); r != nil {
							global.SetError(Replaces(fmt.Sprintf("%v", r)), instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN())
						}
					}()
					if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); !ok {
						instruction.(interfaces.Instruction).Exec(global)
					}
				}()
			}
		}()
	}()

	return ctx.JSON(fiber.Map{
		"console": utils.GetStringOuts(),
	})
}

func (c Controller) GetAST(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ast": TreeDot(ParseTree, ParserG.RuleNames),
	})
}

func (c Controller) GetSymbolsTable(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"table": env.SymTable.GetDot(),
	})
}

func (c Controller) GetErrors(ctx *fiber.Ctx) error {
	errorsDot := `digraph Errors {graph[fontname="Arial" labelloc="t" bgcolor="#252526" fontcolor="white"];node[shape=none fontname="Arial"];label="Errores";table[label=<<table border="0" cellborder="1" cellspacing="0" cellpadding="3"><tr><td bgcolor="#009900" width="100"><font color="#FFFFFF">No.</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Tipo</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Descripción</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Línea</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Columna</font></td></tr>`
	for _, errors := range utils.Errors {
		errorsDot += errors.GetDot()
	}
	errorsDot += `</table>>];}`
	return ctx.JSON(fiber.Map{
		"errors": errorsDot,
	})
}

func Replaces(msg string) string {
	replaces := [][]string{
		{"runtime error: invalid memory address or nil pointer dereference", "Falla en ejeución. Referencia a dirección de memoria inválida o nil"},
		{"interface conversion: interface is nil, not interfaces.Instruction", "No puede ejecutarse la instrucción"},
		{"mismatched input", "Inesperado:"},
		{"expecting", ". Se esperaba:"},
		{"missing", "Inesperado:"},
		{"extraneous input", "Entrada extraña: "},
		{"input", "entrada"},
		{"at", "en"},
		{"no viable alternenive", "Alternativa inviable"},
	}
	for _, m := range replaces {
		msg = strings.Replace(msg, m[0], m[1], -1)
	}
	return msg
}

func TreeDot(tree antlr.Tree, ruleNames []string) string {
	dot := "digraph Tree {"
	dot += "\n\tgraph[fontname=\"Arial\" labelloc=\"t\" bgcolor=\"#252526\"];"
	dot += "\n\tnode[fontname=\"Arial\" fontsize=\"8\" width=\"0\" height=\"0\" color=\"white\" fontcolor=\"white\"];"
	dot += "\n\tedge[fontname=\"Arial\" color=\"white\" dir=none];"
	dot += NodesTree("i", tree, ruleNames)
	dot += "\n}"
	return dot
}

func NodesTree(id string, tree antlr.Tree, ruleNames []string) string {
	s := antlr.TreesGetNodeText(tree, ruleNames, nil)
	s = antlr.EscapeWhitespace(s, false)
	c := tree.GetChildCount()
	res := "\n\tn" + id + "[label=\"" + strings.Replace(s, "\"", "\\\"", -1) + "\"];"
	for i := 0; i < c; i++ {
		res += NodesTree(fmt.Sprintf("%s%v", id, i), tree.GetChild(i), ruleNames)
		res += "\n\tn" + id + " -> " + "n" + fmt.Sprintf("%s%v", id, i) + ";"
	}
	return res
}
