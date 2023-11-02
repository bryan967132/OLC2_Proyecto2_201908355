package utils

var Plus = [][]Type{
	//	     INT    FLOAT  STR
	/*INT*/ {INT, FLOAT, NIL},
	/*FLT*/ {FLOAT, FLOAT, NIL},
	/*STR*/ {NIL, NIL, STRING},
}

var Minus = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var Mult = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var Div = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var Mod = [][]Type{
	//	     INT
	/*INT*/ {INT},
}
