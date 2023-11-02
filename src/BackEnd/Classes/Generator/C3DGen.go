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
	ConcatString    bool
	IntString       bool
	MainC3DCode     bool
	BreakLabel      string
	ContinueLabel   string
}

func NewC3DGen() *C3DGen {
	return &C3DGen{TemporalCount: 0, LabelCount: 0, BreakLabel: "", ContinueLabel: "", PrintString: true, ConcatString: true, IntString: true, MainC3DCode: true}
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

func (g *C3DGen) AddFunc(id string) {
	g.C3DFunctions = append(g.C3DFunctions, "void "+id+"() {")
}

func (g *C3DGen) AddEnd() {
	g.C3DFunctions = append(g.C3DFunctions, "\treturn;")
	g.C3DFunctions = append(g.C3DFunctions, "}\n")
}

func (g *C3DGen) AddCall(id string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t"+id+"();")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t"+id+"();")
	}
}

func (g *C3DGen) AddIf(left string, right string, operator string, Label string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\tif("+left+" "+operator+" "+right+") goto "+Label+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\tif("+left+" "+operator+" "+right+") goto "+Label+";")
	}
}

func (g *C3DGen) AddGoto(Label string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\tgoto "+Label+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\tgoto "+Label+";")
	}
}

func (g *C3DGen) AddExpression(target string, left string, operator string, right string) {
	if g.MainC3DCode {
		if left != "" {
			g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = "+left+" "+operator+" "+right+";")
		} else {
			g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = "+operator+" "+right+";")

		}
	} else {
		if left != "" {
			g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = "+left+" "+operator+" "+right+";")
		} else {
			g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = "+operator+" "+right+";")

		}
	}
}

func (g *C3DGen) AddExpressionInit(target string, value string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = "+value+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = "+value+";")
	}
}

func (g *C3DGen) AddAssign(target, val string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = "+val+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = "+val+";")
	}
}

func (g *C3DGen) AddSetHeap(index string, value string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\theap["+index+"] = "+value+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\theap["+index+"] = "+value+";")
	}
}

func (g *C3DGen) AddGetHeap(target string, index string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = heap["+index+"];")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = heap["+index+"];")
	}
}

func (g *C3DGen) AddSetStack(index string, value string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\tstack["+index+"] = "+value+";")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\tstack["+index+"] = "+value+";")
	}
}

func (g *C3DGen) AddGetStack(target string, index string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t"+target+" = stack["+index+"];")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t"+target+" = stack["+index+"];")
	}
}

func (g *C3DGen) AddPrintf(typePrint string, value string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\tprintf(\"%"+typePrint+"\", "+value+");")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\tprintf(\"%"+typePrint+"\", "+value+");")
	}
}

func (g *C3DGen) AddComment(target string) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "\t/* "+target+" */")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "\t/* "+target+" */")
	}
}

func (g *C3DGen) AddBr() {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, "")
	} else {
		g.C3DFunctions = append(g.C3DFunctions, "")
	}
}

func (g *C3DGen) NewEnv(size int) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, fmt.Sprintf("\tP = P + %v;", size))
	} else {
		g.C3DFunctions = append(g.C3DFunctions, fmt.Sprintf("\tP = P + %v;", size))
	}
}

func (g *C3DGen) PrevEnv(size int) {
	if g.MainC3DCode {
		g.C3DInstructions = append(g.C3DInstructions, fmt.Sprintf("\tP = P - %v;", size))
	} else {
		g.C3DFunctions = append(g.C3DFunctions, fmt.Sprintf("\tP = P - %v;", size))
	}
}

func (g *C3DGen) AddPrint(value string) {
	for _, c := range value {
		g.AddPrintf("c", fmt.Sprintf("(char) %d", byte(c)))
	}
}

func (g *C3DGen) GeneratePrintString() {
	if g.PrintString {
		// Temporales y Etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLbl1 := g.NewLabel()
		newLbl2 := g.NewLabel()
		// Función printString
		g.C3DNatives = append(g.C3DNatives, "void printString() {")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp1+" = P + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = stack[(int) "+newTemp1+"];")
		g.C3DNatives = append(g.C3DNatives, newLbl1+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = heap[(int) "+newTemp2+"];")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp3+" == -1) goto "+newLbl2+";")
		g.C3DNatives = append(g.C3DNatives, "\tprintf(\"%c\", (char) "+newTemp3+");")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = "+newTemp2+" + 1;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLbl1+";")
		g.C3DNatives = append(g.C3DNatives, newLbl2+":")
		g.C3DNatives = append(g.C3DNatives, "\treturn;")
		g.C3DNatives = append(g.C3DNatives, "}\n")
		g.PrintString = false
	}
}

func (g *C3DGen) GenerateConcatString() {
	if g.ConcatString {
		// Temporales y Etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newLbl1 := g.NewLabel()
		newLbl2 := g.NewLabel()
		newLbl3 := g.NewLabel()
		newLbl4 := g.NewLabel()
		// Función concatString
		g.C3DNatives = append(g.C3DNatives, "void concatString() {")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp1+" = H;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = P + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = stack[(int) "+newTemp2+"];")
		g.C3DNatives = append(g.C3DNatives, newLbl1+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp4+" = heap[(int) "+newTemp3+"];")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp4+" == -1) goto "+newLbl2+";")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = "+newTemp4+";")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = "+newTemp3+" + 1;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLbl1+";")
		g.C3DNatives = append(g.C3DNatives, newLbl2+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = P + 2;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = stack[(int) "+newTemp2+"];")
		g.C3DNatives = append(g.C3DNatives, newLbl3+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp4+" = heap[(int) "+newTemp3+"];")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp4+" == -1) goto "+newLbl4+";")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = "+newTemp4+";")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = "+newTemp3+" + 1;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLbl3+";")
		g.C3DNatives = append(g.C3DNatives, newLbl4+":")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = -1;")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, "\tstack[(int) P] = "+newTemp1+";")
		g.C3DNatives = append(g.C3DNatives, "\treturn;")
		g.C3DNatives = append(g.C3DNatives, "}\n")
		g.ConcatString = false
	}
}

func (g *C3DGen) GenerateIntString() {
	if g.IntString {
		// Temporales y Etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newTemp7 := g.NewTemp()
		newLbl1 := g.NewLabel()
		newLbl2 := g.NewLabel()
		newLbl3 := g.NewLabel()
		newLbl4 := g.NewLabel()
		newLbl5 := g.NewLabel()
		// Función intString
		g.C3DNatives = append(g.C3DNatives, "void intString() {")
		g.C3DNatives = append(g.C3DNatives, "\n"+newTemp1+" = H;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp2+" = P + 1;")
		g.C3DNatives = append(g.C3DNatives, "\n"+newTemp3+" = stack[(int) "+newTemp2+"];")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp3+" > 0) goto "+newLbl1+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = -"+newTemp3+";")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = 45;")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, newLbl1+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp4+" = "+newTemp3+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp5+" = "+newTemp3+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp6+" = 0;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp7+" = 1;")
		g.C3DNatives = append(g.C3DNatives, newLbl2+":")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp4+" < 1) goto "+newLbl4+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp7+" = "+newTemp7+" * 10;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp4+" = "+newTemp4+" / 10;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLbl2+";")
		g.C3DNatives = append(g.C3DNatives, newLbl4+":")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp7+" = "+newTemp7+" / 10;")
		g.C3DNatives = append(g.C3DNatives, newLbl3+":")
		g.C3DNatives = append(g.C3DNatives, "\tif("+newTemp7+" < 1) goto "+newLbl5+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp5+" = "+newTemp3+" / "+newTemp7+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp6+" = "+newTemp5+" +48;")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = "+newTemp6+";")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp3+" = (int) "+newTemp3+" % (int) "+newTemp7+";")
		g.C3DNatives = append(g.C3DNatives, "\t"+newTemp7+" = "+newTemp7+" / 10;")
		g.C3DNatives = append(g.C3DNatives, "\tgoto "+newLbl3+";")
		g.C3DNatives = append(g.C3DNatives, newLbl5+":")
		g.C3DNatives = append(g.C3DNatives, "\theap[(int) H] = -1;")
		g.C3DNatives = append(g.C3DNatives, "\tH = H + 1;")
		g.C3DNatives = append(g.C3DNatives, "\nstack[(int) P] = "+newTemp1+";")
		g.C3DNatives = append(g.C3DNatives, "}\n")
		g.IntString = false
	}
}

func (g *C3DGen) GenerateFinalCode() {
	// HEADER
	g.C3DCode = append(g.C3DCode, "/* ------ HEADER ------ */")
	g.C3DCode = append(g.C3DCode, "#include <stdio.h>")
	g.C3DCode = append(g.C3DCode, "")
	g.C3DCode = append(g.C3DCode, "float heap[30101999];")
	g.C3DCode = append(g.C3DCode, "float stack[30101999];")
	g.C3DCode = append(g.C3DCode, "float P;")
	g.C3DCode = append(g.C3DCode, "float H;")
	// TEMPORALS
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		tmpDec := fmt.Sprintf("float %v", tempArr[0])
		for _, s := range tempArr[1:] {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";"
		g.C3DCode = append(g.C3DCode, tmpDec)
	}
	g.C3DCode = append(g.C3DCode, "")
	// NATIVES
	if len(g.C3DNatives) > 0 {
		g.C3DCode = append(g.C3DCode, "/* ------ NATIVES ------ */")
		for _, s := range g.C3DNatives {
			g.C3DCode = append(g.C3DCode, s)
		}
	}
	// FUNCTIONS
	if len(g.C3DFunctions) > 0 {
		g.C3DCode = append(g.C3DCode, "/* ------ FUNCTIONS ------ */")
		for _, s := range g.C3DFunctions {
			g.C3DCode = append(g.C3DCode, s)
		}
	}
	// MAIN
	g.C3DCode = append(g.C3DCode, "/* ------ MAIN ------ */")
	g.C3DCode = append(g.C3DCode, "int main() {")
	g.C3DCode = append(g.C3DCode, "\tP = 0;")
	g.C3DCode = append(g.C3DCode, "\tH = 0;")
	for _, s := range g.C3DInstructions {
		g.C3DCode = append(g.C3DCode, s)
	}
	g.C3DCode = append(g.C3DCode, "\treturn 0;\n}")
}
