package main

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	instructions "TSwift/Classes/Instructions"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "../../../Calificación/Embebidas.swift"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	input := string(fileData)
	inputStream := antlr.NewInputStream(input)
	scanner := parser.NewScanner(inputStream)
	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(tokens)
	parser.RemoveErrorListeners()
	parserErrors := listener.NewTSwfitErrorListener()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.TSwfitListener = listener.NewTSwfitListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// fmt.Println(TreeDot(tree, parser.RuleNames))
	/*archivo, _ := os.Create("../../../Inputs/Tree.dot")
	defer archivo.Close() // Cierra el archivo al final de la función

	// Escribe en el archivo
	escritor := bufio.NewWriter(archivo)
	_, _ = escritor.WriteString(TreeDot(tree, parser.RuleNames))*/

	global := env.NewEnv(nil, "Global")
	c3dgen := C3DGen.NewC3DGen()
	/*func() {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(Replaces(fmt.Sprintf("%v", r)), 0, 0)
			}
		}()
	}()
	for _, fail := range parserErrors.Errors {
		global.SetError(Replaces(fail.Msg), fail.Line, fail.Column)
	}*/

	execute := listener.Code
	for _, instruction := range execute {
		/*func() {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(Replaces(fmt.Sprintf("%v", r)), instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN())
			}
		}()*/
		if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); ok {
			instruction.(interfaces.Instruction).Exec(global, c3dgen)
		}
		//}()
	}
	for _, instruction := range execute {
		/*func() {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(Replaces(fmt.Sprintf("%v", r)), instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN())
			}
		}()*/
		if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); !ok {
			instruction.(interfaces.Instruction).Exec(global, c3dgen)
		}
		//}()
	}
	fmt.Println(utils.GetStringOuts())
	c3dgen.GenerateFinalCode()
	file, _ := os.Create("../../Compilado.cpp")
	defer file.Close()
	_, _ = file.WriteString(strings.Join(c3dgen.GetFinalCode(), "\n"))
}

func Replaces(msg string) string {
	replaces := [][]string{
		{"runtime error: invalid memory address or nil pointer dereference", "Falla en ejeución. Referencia a dirección de memoria inválida o nil."},
		{"mismatched input", "Inesperado:"},
		{"expecting", ". Se esperaba:"},
		{"missing", "Inesperado:"},
		{"extraneous input", "Entrada extraña: "},
		{"input", "entrada"},
		{"at", "en"},
		{"no viable alternenive", "Sin recuperar"},
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
