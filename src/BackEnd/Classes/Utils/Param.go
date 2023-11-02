package utils

type Param struct {
	Line        int
	Column      int
	ExternID    string
	ID          string
	IsInout     bool
	IsPrimitive bool
	IsVector    bool
	Type        *AttribsType
}

func NewParam(line, column int, externID, id string, isInout bool, TypeP *AttribsType) *Param {
	if _, ok := TypeP.Value.(Type); ok {
		return &Param{line, column, externID, id, isInout, true, false, TypeP}
	}
	if _, ok := TypeP.Value.(VectorType); ok {
		return &Param{line, column, externID, id, isInout, false, true, TypeP}
	}
	return nil
}
