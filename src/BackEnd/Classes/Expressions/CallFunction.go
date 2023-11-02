package expressions

import (
	env "TSwift/Classes/Env"
	C3DGen "TSwift/Classes/Generator"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"strconv"
)

type CallFunction struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	ID      string
	Args    []utils.Arg
}

func NewCallFunction(line, column int, id string, args []utils.Arg) *CallFunction {
	return &CallFunction{line, column, utils.CALL_FUNC, id, args}
}

func (c *CallFunction) LineN() int {
	return c.Line
}

func (c *CallFunction) ColumnN() int {
	return c.Column
}

func (c *CallFunction) Exec(Env *env.Env, c3dgen *C3DGen.C3DGen) *utils.ReturnValue {
	global := Env.GetGlobal()
	function := (*global.GetFunction(c.ID, c.Line, c.Column)).(interfaces.Function)
	if function != nil {
		if len(function.GetParams()) == len(c.Args) {

			tmp1 := c3dgen.NewTemp()
			c3dgen.AddExpression(tmp1, "P", "+", strconv.Itoa(Env.Size+1))
			for i := 0; i < len(function.GetParams()); i++ {
				if function.GetParams()[i].ExternID == c.Args[i].ExternID {
					if !function.GetParams()[i].IsInout && !c.Args[i].IsInout {
						// POR VALOR
						var value *utils.ReturnValue = c.Args[i].Exp.(interfaces.Expression).Exec(Env, c3dgen)
						c3dgen.AddSetStack("(int) "+tmp1, value.StrValue)
						if i < len(function.GetParams())-1 {
							c3dgen.AddExpression(tmp1, tmp1, "+", "1")
						}
					} else if function.GetParams()[i].IsInout && c.Args[i].IsInout {
						// POR REFERENCIA
					}
				} else {
					global.SetError("Mal uso de nombre externo", c.Args[i].Line, c.Args[i].Column)
				}
			}

			c3dgen.NewEnv(Env.Size)
			c3dgen.AddCall(c.ID)
			c3dgen.AddGetStack(tmp1, "(int) P")
			c3dgen.PrevEnv(Env.Size)

			if function.GetType() != utils.NIL {
				return &utils.ReturnValue{StrValue: tmp1, Type: function.GetType()}
			}
			return nil
		}
		global.SetError("Cantidad errónea de parámetros enviados", c.Line, c.Column)
		return nil
	}
	return nil
}
