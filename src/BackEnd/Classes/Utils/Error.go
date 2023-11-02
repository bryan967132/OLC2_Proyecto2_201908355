package utils

import "fmt"

type Error struct {
	Number      int
	Line        int
	Column      int
	Type        TypeError
	Description string
}

func NewError(line, column int, Type TypeError, description string) *Error {
	return &Error{0, line, column, Type, description}
}

func (e *Error) ToString() string {
	return fmt.Sprintf("→ Error %v, %s. %v:%v", e.Type, e.Description, e.Line, e.Column)
}

func (e *Error) GetDot() string {
	return fmt.Sprintf(`<tr><td bgcolor="white">%v</td><td bgcolor="white">%v</td><td bgcolor="white">%s</td><td bgcolor="white">%v</td><td bgcolor="white">%v</td></tr>`, e.Number, e.Type, e.Description, e.Line, e.Column)
}
