package C3DGen

import "fmt"

type C3DGen struct {
	TemporalCount   int
	LabelCount      int
	C3DInstructions []string
	C3DCode         []string
	C3DNatives      []string
	C3DFunctions    []string
	Temporals       []string
	PrintString     bool
	MainC3DCode     bool
	BreakLabel      string
	ContinueLabel   string
}

func NewC3DGen() *C3DGen {
	return &C3DGen{TemporalCount: 0, LabelCount: 0, BreakLabel: "", ContinueLabel: "", PrintString: true}
}

func (g *C3DGen) GetCode() []string {
	return g.C3DInstructions
}

func (g *C3DGen) GetFinalCode() []string {
	return g.C3DCode
}

func (g *C3DGen) GetTemps() []string {
	return g.Temporals
}

// add break lbl
func (g *C3DGen) AddBreak(lbl string) {
	g.BreakLabel = lbl
}

// add continue lbl
func (g *C3DGen) AddContinue(lbl string) {
	g.ContinueLabel = lbl
}

// Generar un nuevo temporal
func (g *C3DGen) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.TemporalCount)
	g.TemporalCount++
	//Lo guardamos para declararlo
	g.Temporals = append(g.Temporals, temp)
	return temp
}

// Generador de Label
func (g *C3DGen) NewLabel() string {
	temp := g.LabelCount
	g.LabelCount = g.LabelCount + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Añade Label al codigo
func (g *C3DGen) AddLabel(Label string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, Label+":")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, Label+":")
	}
}

func (g *C3DGen) AddIf(left string, right string, operator string, Label string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "if("+left+" "+operator+" "+right+") goto "+Label+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "if("+left+" "+operator+" "+right+") goto "+Label+";")
	}
}

func (g *C3DGen) AddGoto(Label string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "goto "+Label+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "goto "+Label+";")
	}
}

func (g *C3DGen) AddExpression(target string, left string, right string, operator string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, target+" = "+left+" "+operator+" "+right+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, target+" = "+left+" "+operator+" "+right+";")
	}
}

func (g *C3DGen) GeneratePrintString() {
	if g.PrintString {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.C3DNatives = append(g.C3DNatives, "void dbrust_printString() {")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp1+" = P + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];")
		g.C3DNatives = append(g.C3DNatives, "\t"+newLvl1+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = heap[(int)"+newTemp2+"];")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp3+" == -1) goto "+newLvl2+";")
		g.C3DNatives = append(g.C3DNatives, "\tprintf(\"%c\", (char)"+newTemp3+");")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = "+newTemp2+" + 1;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLvl1+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newLvl2+":")
		g.C3DNatives = append(g.C3DNatives, "\treturn;")
		g.C3DNatives = append(g.C3DNatives, "}\n")
		g.PrintString = false
	}
}

func (g *C3DGen) GenerateFinalCode() {
	// HEADER
	g.C3DCode = append(g.C3DCode, "/*------HEADER------*/")
	g.C3DCode = append(g.C3DCode, "#include <stdio.h>")
	g.C3DCode = append(g.C3DCode, "#include <math.h>")
	g.C3DCode = append(g.C3DCode, "double heap[30101999];")
	g.C3DCode = append(g.C3DCode, "double stack[30101999];")
	g.C3DCode = append(g.C3DCode, "double P = 0;")
	g.C3DCode = append(g.C3DCode, "double H = 0;")
	// TEMPORALS
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		g.C3DCode = append(g.C3DCode, "double ")
		tmpDec := fmt.Sprintf("%v", tempArr[0])
		tempArr = tempArr[1:]
		for _, s := range tempArr {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";"
		g.C3DCode = append(g.C3DCode, tmpDec)
	}
	// NATIVES
	if len(g.C3DNatives) > 0 {
		g.C3DCode = append(g.C3DCode, "/*------NATIVES------*/")
		for _, s := range g.C3DNatives {
			g.C3DCode = append(g.C3DCode, s)
		}
	}
	// FUNCTIONS
	if len(g.C3DFunctions) > 0 {
		g.C3DCode = append(g.C3DCode, "/*------FUNCTIONS------*/")
		for _, s := range g.C3DFunctions {
			g.C3DCode = append(g.C3DCode, s)
		}
	}
	// MAIN
	g.C3DCode = append(g.C3DCode, "/*------MAIN------*/")
	g.C3DCode = append(g.C3DCode, "int main() {")
	for _, s := range g.C3DFunctions {
		g.C3DCode = append(g.C3DCode, s)
	}
	g.C3DCode = append(g.C3DCode, "\treturn 0;\n}")
}
