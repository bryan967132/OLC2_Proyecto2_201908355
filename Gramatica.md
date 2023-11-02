## Expresiones Regulares
| Expresión Regular | Nombre del Token | | Expresión Regular | Nombre del Token |
| - | - | - | - | - |
| ```\s+```                                |                      |  |  ```[ \n\r]```                             |                    |
| ```'//'(~[\r\n])*```                     |                      |  |  ```[/][*]~[*]*[*]+(~[*/]~[*]*[*]+)*[/]``` |                    |
| ```Int```                                | ```RW_Int```         |  | ```Float```                                | ```RW_Float```     |
| ```String```                             | ```RW_String```      |  | ```Bool```                                 | ```RW_Bool```      |
| ```Character```                          | ```RW_Character```   |  | ```var```                                  | ```RW_var```       |
| ```let```                                | ```RW_let```         |  | ```if```                                   | ```RW_if```        |
| ```else```                               | ```RW_else```        |  | ```for```                                  | ```RW_for```       |
| ```while```                              | ```RW_while```       |  | ```guard```                                | ```RW_guard```     |
| ```switch```                             | ```RW_switch```      |  | ```case```                                 | ```RW_case```      |
| ```default```                            | ```RW_default```     |  | ```break```                                | ```RW_break```     |
| ```continue```                           | ```RW_continue```    |  | ```return```                               | ```RW_return```    |
| ```true```                               | ```RW_true```        |  | ```false```                                | ```RW_false```     |
| ```nil```                                | ```RW_nil```         |  | ```func```                                 | ```RW_func```      |
| ```inout```                              | ```RW_inout```       |  | ```in```                                   | ```RW_in```        |
| ```append```                             | ```RW_append```      |  | ```removeLast```                           | ```RW_removeLast```|
| ```remove```                             | ```RW_remove```      |  | ```at```                                   | ```RW_at```        |
| ```isEmpty```                            | ```RW_isEmpty```     |  | ```count```                                | ```RW_count```     |
| ```repeating```                          | ```RW_repeating```   |  | ```struct```                               | ```RW_struct```    |
| ```mutating```                           | ```RW_mutating```    |  | ```self```                                 | ```RW_self```      |
| ```print```                              | ```RW_print```       |  | ```->```                                   | ```TK_prompt```    |
| ```_```                                  | ```TK_under```       |  | ```'"'(\~('\n'\|'"'\|'\\')\|'\\'.)'"'```   | ```TK_char```      |
| ```'"'(~('\n'\|'"'\|'\\')\|'\\'.)*'"'``` | ```TK_string```      |  | ```[0-9]+```                               | ```TK_int```       |
| ```[0-9]+'.'[0-9]+```                    | ```TK_float```       |  | ```('_')*[a-zA-Z][a-zA-Z0-9_]*```          | ```TK_id```        |
| ```+=```                                 | ```TK_add```         |  | ```-=```                                   | ```TK_sub```       |
| ```+```                                  | ```TK_plus```        |  | ```-```                                    | ```TK_minus```     |
| ```*```                                  | ```TK_mult```        |  | ```/```                                    | ```TK_div```       |
| ```%```                                  | ```TK_mod```         |  | ```==```                                   | ```TK_equequ```    |
| ```!=```                                 | ```TK_notequ```      |  | ```<=```                                   | ```TK_lessequ```   |
| ```>=```                                 | ```TK_moreequ```     |  | ```=```                                    | ```TK_equ```       |
| ```<```                                  | ```TK_less```        |  | ```>```                                    | ```TK_more```      |
| ```&&```                                 | ```TK_and```         |  | ```||```                                   | ```TK_or```        |
| ```!```                                  | ```TK_not```         |  | ```(```                                    | ```TK_lpar```      |
| ```)```                                  | ```TK_rpar```        |  | ```{```                                    | ```TK_lbrc```      |
| ```}```                                  | ```TK_rbrc```        |  | ```[```                                    | ```TK_lbrk```      |
| ```]```                                  | ```TK_rbrk```        |  | ```.```                                    | ```TK_dot```       |
| ```,```                                  | ```TK_comma```       |  | ```:```                                    | ```TK_colon```     |
| ```;```                                  | ```TK_semicolon```   |  | ```?```                                    | ```TK_question```  |
| ```&```                                  | ```TK_amp```         |  |                                            |                    |

## Terminales
| Terminal | Terminal | Terminal | Terminal | Terminal | Terminal |
| - | - | - | - | - | - |
| ```RW_Int```          | ```RW_Float```       | ```RW_Bool```        | ```RW_Character```   | ```RW_String```      | ```RW_true```        |
| ```RW_false```        | ```RW_var```         | ```RW_let```         | ```RW_if```          | ```RW_else```        | ```RW_for```         |
| ```RW_while```        | ```RW_guard```       | ```RW_switch```      | ```RW_case```        | ```RW_default```     | ```RW_break```       |
| ```RW_continue```     | ```RW_return```      | ```RW_nil```         | ```RW_func```        | ```RW_inout```       | ```RW_in```          |
| ```RW_append```       | ```RW_removeLast```  | ```RW_remove```      | ```RW_at```          | ```RW_isEmpty```     | ```RW_count```       |
| ```RW_repeating```    | ```RW_struct```      | ```RW_mutating```    | ```RW_self```        | ```RW_print```       | ```TK_prompt```      |
| ```TK_under```        | ```TK_char```        | ```TK_string```      | ```TK_int```         | ```TK_float```       | ```TK_id```          |
| ```TK_add```          | ```TK_sub```         | ```TK_plus```        | ```TK_minus```       | ```TK_mult```        | ```TK_div```         |
| ```TK_mod```          | ```TK_equequ```      | ```TK_notequ```      | ```TK_lessequ```     | ```TK_moreequ```     | ```TK_equ```         |
| ```TK_less```         | ```TK_more```        | ```TK_and```         | ```TK_or```          | ```TK_not```         | ```TK_lpar```        |
| ```TK_rpar```         | ```TK_lbrc```        | ```TK_rbrc```        | ```TK_lbrk```        | ```TK_rbrk```        | ```TK_dot```         |
| ```TK_comma```        | ```TK_colon```       | ```TK_semicolon```   | ```TK_question```    | ```TK_amp```         |

## No Terminales
| No Terminal | No Terminal | No Terminal | No Terminal |
| - | - | - | - |
| ```init```             | ```instsglobal```        | ```instglobal```          | ```callfunc```            |
| ```listargs```         | ```arg```                | ```decvar```              | ```deccst```              |
| ```declfunc```         | ```listparams```         | ```param```               | ```typeparam```           |
| ```ifstruct```         | ```switchstruct```       | ```envs```                | ```casesdefault```        |
| ```cases```            | ```case```               | ```default```             | ```loopfor```             |
| ```range```            | ```loopwhile```          | ```guard```               | ```reasign```             |
| ```addsub```           | ```decvector```          | ```funcvector```          | ```decmatrix```           |
| ```defvector```        | ```defmatrix```          | ```matrix```              | ```vectors```             |
| ```vector```           | ```simplematrix```       | ```typematrix```          | ```reasignvector```       |
| ```dims```             | ```defstruct```          | ```listattribs```         | ```attrib```              |
| ```useattribs```       | ```useattribs1```        | ```obj```                 | ```print```               |
| ```env```              | ```instructions```       | ```instruction```         | ```type```                |
| ```typeComp```         | ```exp```                |

## Producciones
```antlr4
init :
	instsglobal EOF |
	EOF

instsglobal :
	instsglobal instglobal |
	instglobal             ;

instglobal :
	instruction |
	declfunc    |
	defstruct   ;

callfunc :
	TK_id '(' listargs ')' |
	TK_id '(' ')'          ;

listargs :
	listargs ',' arg |
	arg              ;

arg :
	TK_id ':' '&' exp |
	TK_id ':' exp     |
	'&' exp           |
	exp               ;

decvar :
	'var' TK_id ':' type '=' exp |
	'var' TK_id ':' type '?'     |
	'var' TK_id '=' exp          ;

deccst :
	'let' TK_id ':' type '=' exp |
	'let' TK_id '=' exp          ;

declfunc :
	'func' TK_id '(' listparams ')' '->' typeComp env |
	'func' TK_id '(' listparams ')' env               |
	'func' TK_id '(' ')' '->' typeComp env            |
	'func' TK_id '(' ')' env                          ;

listparams :
	listparams ',' param |
	param                ;

param :
	TK_id TK_id ':' 'inout' typeparam |
    TK_id       ':' 'inout' typeparam |
    '_'   TK_id ':' 'inout' typeparam |
    TK_id       ':' 'inout' typeparam |
    TK_id TK_id ':'         typeparam |
    TK_id       ':'         typeparam |
    '_'   TK_id ':'         typeparam |
    TK_id       ':'         typeparam ;

typeparam :
	typeComp   |
	typematrix ;

ifstruct :
	'if' exp env 'else' ifstruct |
	'if' exp env 'else' env      |
	'if' exp env                 ;

switchstruct :
	'switch' exp envs ;

envs :
	'{' casesdefault '}' |
	'{' '}'              ;

casesdefault :
	cases default |
	cases         |
	default       ;

cases :
	cases case |
	case       ;

case :
	'case' exp ':' instructions |
	'case' exp ':'              ;

default :
	'default' ':' instructions |
	'default' ':'              ;

loopfor :
	'for' ('_' | TK_id) 'in' range env |
	'for' ('_' | TK_id) 'in' exp env   ;

range :
	exp '.' '.' '.' exp ;

loopwhile :
	'while' exp env ;

guard :
	'guard' exp 'else' env ;

reasign :
	TK_id '=' exp ;

addsub :
	TK_id ('+=' | '-=') exp ;

decvector :
	'var' TK_id ':' '[' typeComp ']' '=' defvector |
	'let' TK_id ':' '[' typeComp ']' '=' defvector |
	'var' TK_id '=' defvector                      |
	'let' TK_id '=' defvector                      ;

defvector :
	'[' listexp ']'          |
	'[' ']'                  |
	'[' typeComp ']' '(' ')' |
	TK_id                    ;

listexp :
	listexp ',' exp |
	exp             ;

funcvector :
	TK_id '.' 'append' '(' exp ')'          |
	TK_id '.' 'removeLast' '(' ')'          |
	TK_id '.' 'remove' '(' 'at' ':' exp ')' ;

decmatrix :
	'var' TK_id ':' typematrix '=' defmatrix |
	'var' TK_id '=' defmatrix                ;

defmatrix :
	matrix       |
	simplematrix ;

matrix :
	'[' vectors ']' ;

vectors :
	vectors ',' vector |
	vector             ;

vector :
	'[' listexp ']' |
	matrix          ;

simplematrix :
	typematrix '(' 'repeating' ':' simplematrix ',' 'count' ':' exp ')' |
	typematrix '(' 'repeating' ':' exp ',' 'count' ':' exp ')'          ;

typematrix :
	'[' typematrix ']' |
	'[' type ']'       ;

reasignvector :
	TK_id dims '=' exp ;

dims :
	dims '[' exp ']' |
	'[' exp ']'      ;

defstruct :
	'struct' TK_id '{' listattribs '}' ;

listattribs :
	listattribs ';'? attrib |
	attrib ';'?             ;

attrib :
	('let' | 'var') TK_id (':' typeComp)? ('=' exp)? |
	'mutating'? declfunc                             ;

decstruct :
	('let' | 'var') TK_id (':' TK_id)? '=' TK_id '(' listdupla? ')' |
	('let' | 'var') TK_id (':' TK_id)? '=' TK_id '(' ')'            ;

listdupla :
	TK_id ':' exp ',' listdupla |
	TK_id ':' exp               ;

useattribs :
	obj useattribs1  |
	obj '.' callfunc ;

obj :
	TK_id '[' exp ']' |
	TK_id             ;

useattribs1 :
	'.' TK_id useattribs1 |
	'.' TK_id             ;

print :
	'print' '(' listexp ')' |
	'print' '(' ')'         ;

env :
	'{' instructions '}' |
	'{' '}'              ;

instructions :
	instructions instruction |
	instruction              ;

instruction :
	decvar                       ';'? |
	deccst                       ';'? |
	ifstruct                          |
	switchstruct                      |
	loopfor                           |
	loopwhile                         |
	guard                             |
	('self' '.' )? reasign       ';'? |
	('self' '.' )? addsub        ';'? |
	decvector                    ';'? |
	funcvector                   ';'? |
	('self' '.' )? reasignvector ';'? |
	decmatrix                    ';'? |
	decstruct                    ';'? |
	('self' '.' )? useattribs    ';'? |
	('self' '.' )? callfunc      ';'? |
	print                        ';'? |
	'return' exp                 ';'? |
	'return'                     ';'? |
	'continue'                   ';'? |
	'break'                      ';'? ;

type :
	'String'    |
	'Int'       |
	'Bool'      |
	'Character' |
	'Float'     ;

typeComp :
	type  |
	TK_id ;

exp :
	TK_id dims                |
	exp '.' 'isEmpty'         |
	exp '.' 'count'           |
	'-' exp                   |
	exp ('*' | '/' | '%') exp |
	exp ('+' | '-') exp       |
	exp ('<=' | '>=') exp     |
	exp ('<'  | '>') exp      |
	exp ('==' | '!=') exp     |
	'!' exp                   |
	exp '&&' exp              |
	exp '||' exp              |
	type '(' exp ')'          |
	('self' '.' )? useattribs |
	('self' '.' )? callfunc   |
	('self' '.' )? TK_id      |
	'nil'                     |
	TK_string                 |
	TK_char                   |
	TK_int                    |
	TK_float                  |
	'true'                    |
	'false'                   |
	'(' exp ')'               ;
```