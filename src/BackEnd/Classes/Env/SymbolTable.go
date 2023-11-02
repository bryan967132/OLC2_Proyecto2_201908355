package env

type SymbolTable struct {
	Symbols []SymTab
}

func NewSymbolTable() SymbolTable {
	return SymbolTable{[]SymTab{}}
}

func (t *SymbolTable) Push(sym SymTab) {
	if t.validateSymbol(sym) {
		t.Symbols = append(t.Symbols, sym)
	}
}

func (t *SymbolTable) validateSymbol(sym SymTab) bool {
	for _, i := range t.Symbols {
		if i.Hash() == sym.Hash() {
			return false
		}
	}
	return true
}

func (t *SymbolTable) GetDot() string {
	dot := `digraph SymbolsTable {graph[fontname="Arial" labelloc="t" bgcolor="#252526" fontcolor="white"];node[shape=none fontname="Arial"];label="Tabla de Símbolos";table[label=<<table border="0" cellborder="1" cellspacing="0" cellpadding="3"><tr><td bgcolor="#009900" width="100"><font color="#FFFFFF">No.</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Identificador</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Tipo</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Tipo de Dato</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Entorno</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Línea</font></td><td bgcolor="#009900" width="100"><font color="#FFFFFF">Columna</font></td></tr>`
	for i, sym := range t.Symbols {
		sym.Num = i + 1
		dot += sym.GetDot()
	}
	dot += `</table>>];}`
	return dot
}

var SymTable SymbolTable = NewSymbolTable()
