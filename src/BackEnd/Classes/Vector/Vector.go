package vector

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Vector struct {
	IsMatrix    bool
	IsReuseID   bool
	IsRepeating bool
	Length      int
	Dims        int
	Type        *utils.AttribsType
	Vectors     []*Vector
	Elements    []interfaces.Expression
	Values      []*utils.ReturnType
	ID          string
	Repeating   *Repeating
}

func NewMatrix(Type *utils.AttribsType, Vectors []*Vector) *Vector {
	return &Vector{IsMatrix: true, IsReuseID: false, IsRepeating: false, Type: Type, Vectors: Vectors}
}

func NewVector(Type *utils.AttribsType, Elements []interfaces.Expression) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: false, IsRepeating: false, Type: Type, Elements: Elements}
}

func NewMatrixRepeating(repeating *Repeating) *Vector {
	return &Vector{IsMatrix: true, IsRepeating: true, Repeating: repeating}
}

func NewReuseVector(id string) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: true, IsRepeating: false, ID: id}
}

func (v *Vector) Generate(env *env.Env, Type utils.Type) bool {
	if v.IsMatrix {
		for i := 0; i < len(v.Vectors); i++ {
			v.Vectors[i].Generate(env, Type)
		}
		v.Length = len(v.Vectors)
		v.Dims = v.Vectors[0].Dims + 1
		if v.Type == nil {
			v.Type = v.Vectors[0].Type
		}
		return true
	}
	v.Values = []*utils.ReturnType{}
	for i := 0; i < len(v.Elements); i++ {
		v.Values = append(v.Values, v.Elements[i].Exec(env))
		if Type == utils.NIL {
			Type = v.Values[i].Type
		}
		if v.Values[i].Type != Type {
			if v.Values[i].Type == utils.CHAR && Type == utils.STRING || v.Values[i].Type == utils.INT && Type == utils.FLOAT {
				v.Values[i].Type = Type
				continue
			}
			return false
		}
	}
	v.Length = len(v.Elements)
	v.Dims = 1
	if v.Type == nil {
		v.Type = utils.NewAttribsType(0, 0, Type, true)
	}
	return true
}

func (v *Vector) GenerateRepeating(env *env.Env, line, column int, Type utils.Type) *Vector {
	n := v.generateRepeating(env, line, column, Type, v.Repeating)
	if n != nil {
		return n.GetCopy()
	}
	return nil
}

func (v *Vector) generateRepeating(env *env.Env, line, column int, Type utils.Type, repeating *Repeating) *Vector {
	if repeating.Repeating != nil {
		vec := v.generateRepeating(env, line, column, Type, repeating.Repeating)
		if vec != nil {
			if vec.Type.Value.(utils.Type) == repeating.Type {
				count := repeating.Times.Exec(env)
				if count.Type == utils.INT {
					vectors := []*Vector{}
					for i := 0; i < count.Value.(int); i++ {
						vectors = append(vectors, vec.GetCopy())
					}
					mat := NewMatrix(utils.NewAttribsType(0, 0, Type, true), vectors)
					mat.Dims = repeating.Dims
					mat.Length = len(vectors)
					return mat
				}
			}
			env.SetError("Los tipos no coinciden para el vector", line, column)
		}
		return nil
	}
	count := repeating.Times.Exec(env)
	if count.Type == utils.INT {
		elements := []interfaces.Expression{}
		values := []*utils.ReturnType{}
		var value *utils.ReturnType
		value = repeating.Value.Exec(env)
		if Type == utils.NIL {
			Type = value.Type
		}
		if value.Type != repeating.Type {
			if value.Type == utils.INT && repeating.Type == utils.FLOAT || value.Type == utils.CHAR && repeating.Type == utils.STRING {
				value.Type = repeating.Type
			} else {
				env.SetError("Los tipos no coinciden para el vector", line, column)
				return nil
			}
		}
		if Type != value.Type {
			env.SetError("Los tipos no coinciden para el vector", line, column)
			return nil
		}
		for i := 0; i < count.Value.(int); i++ {
			elements = append(elements, repeating.Value)
			values = append(values, value)
		}
		vec := NewVector(utils.NewAttribsType(0, 0, Type, true), elements)
		vec.Dims = repeating.Dims
		vec.Length = len(elements)
		vec.Values = values
		return vec
	}
	env.SetError("Los tipos no coinciden para el vector", line, column)
	return nil
}

func (v *Vector) GetPosition(env *env.Env, indexs []int, line, column int) interface{} {
	if len(indexs) > v.Dims {
		env.SetError("Las dimensiones no coinciden con las del vector", line, column)
		return nil
	}
	if len(indexs) == 1 {
		if v.IsMatrix {
			if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
				return v.Vectors[indexs[0]]
			}
			env.SetError("Índices fuera de rango", line, column)
			return nil
		}
		if indexs[0] >= 0 && indexs[0] < len(v.Values) {
			return v.Values[indexs[0]]
		}
		env.SetError("Índices fuera de rango", line, column)
		return nil
	}
	if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
		return v.Vectors[indexs[0]].GetPosition(env, indexs[1:], line, column)
	}
	env.SetError("Índices fuera de rango", line, column)
	return nil
}

func (v *Vector) SetValuePosition(env *env.Env, indexs []int, newValue interfaces.Expression, line, column int) bool {
	if len(indexs) > v.Dims {
		env.SetError("Las dimensiones no coinciden con las del vector", line, column)
		return false
	}
	if len(indexs) == 1 {
		if v.IsMatrix {
			if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
				env.SetError("Solo puede modificarse un elemento a la vez en un vector", line, column)
				return false
			}
			env.SetError("Índices fuera de rango", line, column)
			return false
		}
		if indexs[0] >= 0 && indexs[0] < len(v.Values) {
			if v.Values[indexs[0]].Type == newValue.Exec(env).Type {
				v.Elements[indexs[0]] = newValue
				v.Values[indexs[0]] = newValue.Exec(env)
				return true
			}
			env.SetError("El tipo del nuevo valor no coincide con el que almacena el vector", line, column)
			return false
		}
		env.SetError("Índices fuera de rango", line, column)
		return false
	}
	if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
		return v.Vectors[indexs[0]].SetValuePosition(env, indexs[1:], newValue, line, column)
	}
	env.SetError("Índices fuera de rango", line, column)
	return false
}

func (v *Vector) GetCopy() *Vector {
	vector := &Vector{}
	vector.IsMatrix = v.IsMatrix
	vector.IsReuseID = false
	vector.Length = v.Length
	vector.Dims = v.Dims
	vector.Type = v.Type
	vector.ID = v.ID
	for _, v := range v.Vectors {
		vector.Vectors = append(vector.Vectors, v.GetCopy())
	}
	for _, v := range v.Elements {
		vector.Elements = append(vector.Elements, v)
	}
	for _, v := range v.Values {
		vector.Values = append(vector.Values, v)
	}
	return vector
}
