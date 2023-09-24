package utils

var PrintConsole = []string{}
var Errors = []Error{}

func GetStringOuts() string {
	out := ""
	for i := 0; i < len(PrintConsole); i++ {
		out += PrintConsole[i]
		if i < len(PrintConsole)-1 {
			out += "\n"
		}
	}
	if len(Errors) > 0 {
		if out != "" {
			out += "\n\n↳ ERRORES\n"
		} else {
			out += "↳ ERRORES\n"
		}
		for i := 0; i < len(Errors); i++ {
			Errors[i].Number = i + 1
			out += Errors[i].ToString()
			if i < len(Errors)-1 {
				out += "\n"
			}
		}
	}
	return out
}
