package utils

type TypeInst string

const (
	PRINT            TypeInst = "PRINT"
	INIT_ID          TypeInst = "INIT_ID"
	ASIGN_ID         TypeInst = "ASIGN_ID"
	INIT_FUNCTION    TypeInst = "INIT_FUNCTION"
	MAIN             TypeInst = "MAIN"
	BLOCK_INST       TypeInst = "BLOCK_INST"
	IF               TypeInst = "IF"
	LOOP_FOR         TypeInst = "LOOP_FOR"
	LOOP_WHILE       TypeInst = "LOOP_WHILE"
	GUARD            TypeInst = "GUARD"
	SWITCH           TypeInst = "SWITCH"
	CASE             TypeInst = "CASE"
	BREAK            TypeInst = "BREAK"
	CONTINUE         TypeInst = "CONTINUE"
	ADD              TypeInst = "ADD"
	SUB              TypeInst = "SUB"
	ARRAY_REMOVE     TypeInst = "ARRAY_REMOVE"
	ARRAY_REMOVELAST TypeInst = "ARRAY_REMOVELAST"
	ARRAY_APPEND     TypeInst = "ARRAY_APPEND"
	ASIGN_ARRAY      TypeInst = "ASIGN_ARRAY"
)
