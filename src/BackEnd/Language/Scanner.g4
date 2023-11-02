lexer grammar Scanner;

options {
    language = Go;
}

fragment UNUSED   : [ \r\t]+                            ;
fragment CONTENT  : (~('\n'|'"'|'\\')|'\\'.)            ;
fragment ID       : ('_')*[a-zA-Z][a-zA-Z0-9_]*         ;
fragment CHAR     : '"'(CONTENT)'"'                     ;
fragment STRING   : '"'(CONTENT)*'"'                    ;
fragment INTEGER  : [0-9]+                              ;
fragment FLOAT    : [0-9]+'.'[0-9]+                     ;
fragment COMMENTS : '//'(~[\r\n])*                      ;
fragment COMMENTM : [/][*]~[*]*[*]+(~[*/]~[*]*[*]+)*[/] ;

// Reservadas
RW_Int        : 'Int'         ;
RW_Float      : 'Float'       ;
RW_String     : 'String'      ;
RW_Bool       : 'Bool'        ;
RW_Character  : 'Character'   ;
RW_var        : 'var'         ;
RW_let        : 'let'         ;
RW_if         : 'if'          ;
RW_else       : 'else'        ;
RW_for        : 'for'         ;
RW_while      : 'while'       ;
RW_guard      : 'guard'       ;
RW_switch     : 'switch'      ;
RW_case       : 'case'        ;
RW_default    : 'default'     ;
RW_break      : 'break'       ;
RW_continue   : 'continue'    ;
RW_return     : 'return'      ;
RW_true       : 'true'        ;
RW_false      : 'false'       ;
RW_nil        : 'nil'         ;
RW_func       : 'func'        ;
RW_inout      : 'inout'       ;
RW_in         : 'in'          ;
RW_append     : 'append'      ;
RW_removeLast : 'removeLast'  ;
RW_remove     : 'remove'      ;
RW_at         : 'at'          ;
RW_isEmpty    : 'isEmpty'     ;
RW_count      : 'count'       ;
RW_repeating  : 'repeating'   ;
RW_struct     : 'struct'      ;
RW_mutating   : 'mutating'    ;
RW_self       : 'self'        ;
RW_print      : 'print'       ;
// Operators
TK_prompt     : '->'          ;
TK_under      : '_'           ;
// Valores
TK_char       : CHAR          ;
TK_string     : STRING        ;
TK_int        : INTEGER       ;
TK_float      : FLOAT         ;
// Variables
TK_id         : ID            ;
// Concatenacion
TK_add        : '+='          ;
TK_sub        : '-='          ;
// Operadores Aritméticos
TK_plus       : '+'           ;
TK_minus      : '-'           ;
TK_mult       : '*'           ;
TK_div        : '/'           ;
TK_mod        : '%'           ;
// Operadores Relacionales
TK_equequ     : '=='          ;
TK_notequ     : '!='          ;
TK_lessequ    : '<='          ;
TK_moreequ    : '>='          ;
TK_equ        : '='           ;
TK_less       : '<'           ;
TK_more       : '>'           ;
// Operadores Lógicos
TK_and        : '&&'          ;
TK_or         : '||'          ;
TK_not        : '!'           ;
// Símbolos de Agrupación
TK_lpar       : '('           ;
TK_rpar       : ')'           ;
TK_lbrc       : '{'           ;
TK_rbrc       : '}'           ;
TK_lbrk       : '['           ;
TK_rbrk       : ']'           ;
// Signos
TK_dot        : '.'           ;
TK_comma      : ','           ;
TK_colon      : ':'           ;
TK_semicolon  : ';'           ;
TK_question   : '?'           ;
TK_amp        : '&'           ;
//
NEWLINE       : '\n'     -> skip ;
UNUSED_       : UNUSED   -> skip ;
COMMENTS_     : COMMENTS -> skip ;
COMMENTM_     : COMMENTM -> skip ;