grammar Parser;

@header {
    import (
        expressions "TSwift/Classes/Expressions"
        instructions "TSwift/Classes/Instructions"
        interfaces "TSwift/Classes/Interfaces"
        utils "TSwift/Classes/Utils"
        vector "TSwift/Classes/Vector"
    )
}

options {
    language = Go;
    tokenVocab = Scanner;
}

init returns[[]interface{} result]:
    insts = instsglobal EOF {$result = $insts.result  } |
    EOF                     {$result = []interface{}{}} ;

instsglobal returns[[]interface{} result]:
    l = instsglobal i = instglobal {$result = $l.result;; $result = append($result, $i.result)} |
    i = instglobal                 {$result = []interface{}{$i.result}                        } ;

instglobal returns[interface{} result] :
    inst0 = instruction             {$result = $inst0.result} |
    inst1 = declfunc                {$result = $inst1.result} |
    inst2 = defstruct TK_semicolon? {$result = $inst2.result} ;

callfunc returns[interfaces.Expression result] :
    id = TK_id TK_lpar l = listargs TK_rpar {$result = expressions.NewCallFunction($id.line, $id.pos, $id.text, $l.result)    } |
    id = TK_id TK_lpar TK_rpar              {$result = expressions.NewCallFunction($id.line, $id.pos, $id.text, []utils.Arg{})} ;

listargs returns[[]utils.Arg result] :
    l = listargs TK_comma a = arg {$result = $l.result;; $result = append($result, $a.result)} |
    a = arg                       {$result = []utils.Arg{$a.result}                          } ;

arg returns[utils.Arg result] :
    id = TK_id  TK_colon TK_amp ep = exp {$result = *utils.NewArg($id.line, $id.pos, $id.text, $ep.result, true) } |
    id = TK_id  TK_colon        ep = exp {$result = *utils.NewArg($id.line, $id.pos, $id.text, $ep.result, false)} |
    am = TK_amp ep = exp                 {$result = *utils.NewArg($am.line, $am.pos, "_",      $ep.result, true) } |
    ep = exp                             {$result = *utils.NewArg($ep.result.LineN(), $ep.result.ColumnN(), "_", $ep.result, false)} ;

decvar returns[interfaces.Instruction result] :
    d = RW_var id = TK_id TK_colon t = type TK_equ e = exp {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, $t.result.Value.(utils.Type), $e.result)} |
    d = RW_var id = TK_id TK_colon t = type TK_question    {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, $t.result.Value.(utils.Type), nil)      } |
    d = RW_var id = TK_id TK_equ e = exp                   {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, utils.NIL, $e.result)                   } ;

deccst returns[interfaces.Instruction result] :
    d = RW_let id = TK_id TK_colon t = type TK_equ e = exp {$result = instructions.NewInitID($d.line, $d.pos, false, $id.text, $t.result.Value.(utils.Type), $e.result)} |
    d = RW_let id = TK_id TK_equ e = exp                   {$result = instructions.NewInitID($d.line, $d.pos, false, $id.text, utils.NIL, $e.result)                   } ;

declfunc returns[interfaces.Instruction result] :
    f = RW_func id = TK_id TK_lpar l = listparams TK_rpar TK_prompt t = typeComp b = env {$result = instructions.NewFunction($f.line, $f.pos, $id.text, $l.result,       $b.result, $t.result.Value.(utils.Type))} |
    f = RW_func id = TK_id TK_lpar l = listparams TK_rpar b = env                        {$result = instructions.NewFunction($f.line, $f.pos, $id.text, $l.result,       $b.result, utils.NIL)                   } |
    f = RW_func id = TK_id TK_lpar TK_rpar TK_prompt t = typeComp b = env                {$result = instructions.NewFunction($f.line, $f.pos, $id.text, []utils.Param{}, $b.result, $t.result.Value.(utils.Type))} |
    f = RW_func id = TK_id TK_lpar TK_rpar b = env                                       {$result = instructions.NewFunction($f.line, $f.pos, $id.text, []utils.Param{}, $b.result, utils.NIL)                   } ;

listparams returns[[]utils.Param result] :
    l = listparams TK_comma p = param {$result = $l.result;; $result = append($result, $p.result)} |
    p = param                         {$result = []utils.Param{$p.result}                        } ;

param returns[utils.Param result] :
    ext = TK_id    id = TK_id TK_colon RW_inout t = typeparam {$result = *utils.NewParam($ext.line, $ext.pos, $ext.text, $id.text, true,  $t.result)} |
    id  = TK_id               TK_colon RW_inout t = typeparam {$result = *utils.NewParam($id.line,  $id.pos,  $id.text,  $id.text, true,  $t.result)} |
    ext = TK_under id = TK_id TK_colon RW_inout t = typeparam {$result = *utils.NewParam($ext.line, $ext.pos, $ext.text, $id.text, true,  $t.result)} |
    id  = TK_id               TK_colon RW_inout t = typeparam {$result = *utils.NewParam($id.line,  $id.pos,  $id.text,  $id.text, true,  $t.result)} |
    ext = TK_id    id = TK_id TK_colon          t = typeparam {$result = *utils.NewParam($ext.line, $ext.pos, $ext.text, $id.text, false, $t.result)} |
    id  = TK_id               TK_colon          t = typeparam {$result = *utils.NewParam($id.line,  $id.pos,  $id.text,  $id.text, false, $t.result)} |
    ext = TK_under id = TK_id TK_colon          t = typeparam {$result = *utils.NewParam($ext.line, $ext.pos, $ext.text, $id.text, false, $t.result)} |
    id  = TK_id               TK_colon          t = typeparam {$result = *utils.NewParam($id.line,  $id.pos,  $id.text,  $id.text, false, $t.result)} ;

typeparam returns[*utils.AttribsType result] :
    tc = typeComp   {t := $tc.result;; $result = &t} |
    tm = typematrix {$result = utils.NewAttribsType($tm.result.Line, $tm.result.Column, $tm.result, false)} ;

ifstruct returns[interfaces.Instruction result] :
    r = RW_if cn = exp b1 = env RW_else b2 = ifstruct  {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, $b2.result)                                        } |
    r = RW_if cn = exp b1 = env RW_else b3 = env       {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, (interface{}($b3.result)).(interfaces.Instruction))} |
    r = RW_if cn = exp b1 = env                        {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, nil)                                               } ;

switchstruct returns[interfaces.Instruction result] :
    s = RW_switch e = exp b = envs {$result = instructions.NewSwitch($s.line, $s.pos, $e.result, $b.result[0], $b.result[1].(interfaces.Instruction))} ;

envs returns[[]interface{} result] :
    TK_lbrc cd = casesdefault TK_rbrc {$result = $cd.result             } |
    TK_lbrc TK_rbrc                   {$result = []interface{}{nil, nil}} ;

casesdefault returns[[]interface{} result] :
    c = cases d = default {$result = []interface{}{$c.result, $d.result}} |
    c = cases             {$result = []interface{}{$c.result, nil}      } |
    d = default           {$result = []interface{}{nil,       $d.result}} ;

cases returns[[]interfaces.Instruction result] :
    l = cases c = case {$result = $l.result;; $result = append($result, $c.result)} |
    c = case           {$result = []interfaces.Instruction{$c.result}             } ;

case returns[interfaces.Instruction result] :
    c = RW_case e = exp TK_colon b = instructions {$result = instructions.NewCase($c.line, $c.pos, $e.result, instructions.NewBlock($c.line, $c.pos, $b.result))      } |
    c = RW_case e = exp TK_colon                  {$result = instructions.NewCase($c.line, $c.pos, $e.result, instructions.NewBlock($c.line, $c.pos, []interface{}{}))} ;

default returns[interfaces.Instruction result] :
    d = RW_default TK_colon b = instructions {$result = instructions.NewBlock($d.line, $d.pos, $b.result)      } |
    d = RW_default TK_colon                  {$result = instructions.NewBlock($d.line, $d.pos, []interface{}{})} ;

loopfor returns[interfaces.Instruction result] :
    f = RW_for id = (TK_id | TK_under) RW_in r = range b = env {$result = instructions.NewFor($f.line, $f.pos, $id.text, nil, $r.result[0], $r.result[1], $b.result)} |
    f = RW_for id = (TK_id | TK_under) RW_in e = exp b = env   {$result = instructions.NewFor($f.line, $f.pos, $id.text, $e.result, nil, nil, $b.result)            } ;

range returns[[]interfaces.Expression result] :
    e1 = exp TK_dot TK_dot TK_dot e2 = exp {$result = []interfaces.Expression{$e1.result, $e2.result}} ;

loopwhile returns[interfaces.Instruction result] :
    w = RW_while e = exp b = env {$result = instructions.NewWhile($w.line, $w.pos, $e.result, $b.result)} ;

guard returns[interfaces.Instruction result] :
    g = RW_guard e = exp RW_else b = env {$result = instructions.NewGuard($g.line, $g.pos, $e.result, $b.result)} ;

reasign returns[interfaces.Instruction result] :
    id = TK_id TK_equ e = exp {$result = instructions.NewAsignID($id.line, $id.pos, $id.text, $e.result)} ;

addsub returns[interfaces.Instruction result] :
    id = TK_id s = (TK_add | TK_sub) e = exp {$result = instructions.NewAddSub($id.line, $id.pos, $id.text, $s.text, $e.result)} ;

decvector returns[interfaces.Instruction result] :
    d = RW_var id = TK_id TK_colon TK_lbrk t = typeComp TK_rbrk TK_equ df = defvector {attType := $t.result;; $result = instructions.NewInitVector($d.line, $d.pos, true,  $id.text, &attType, $df.result);} |
    d = RW_let id = TK_id TK_colon TK_lbrk t = typeComp TK_rbrk TK_equ df = defvector {attType := $t.result;; $result = instructions.NewInitVector($d.line, $d.pos, false, $id.text, &attType, $df.result);} |
    d = RW_var id = TK_id TK_equ df = defvector {$result = instructions.NewInitVector($d.line, $d.pos, true,  $id.text, nil, $df.result);} |
    d = RW_let id = TK_id TK_equ df = defvector {$result = instructions.NewInitVector($d.line, $d.pos, false, $id.text, nil, $df.result);} ;

defvector returns[*vector.Vector result] :
    lb = TK_lbrk l = listexp TK_rbrk                  {$result = vector.NewVector(nil, $l.result)                                            } |
    lb = TK_lbrk TK_rbrk                              {$result = vector.NewVector(nil, []interfaces.Expression{})                            } |
    lb = TK_lbrk t = typeComp TK_rbrk TK_lpar TK_rpar {attType := $t.result;; $result = vector.NewVector(&attType, []interfaces.Expression{})} |
    id = TK_id                                        {$result = vector.NewReuseVector($id.text)                                             } ;

listexp returns[[]interfaces.Expression result] :
    l = listexp TK_comma e = exp {$result = $l.result;; $result = append($result, $e.result)} |
    e = exp                      {$result = []interfaces.Expression{$e.result}              } ;

funcvector returns[interfaces.Instruction result] :
    id = TK_id TK_dot RW_append TK_lpar e = exp TK_rpar                {$result = instructions.NewAppend($id.line, $id.pos, $id.text, $e.result)} |
    id = TK_id TK_dot RW_removeLast TK_lpar TK_rpar                    {$result = instructions.NewRemoveLast($id.line, $id.pos, $id.text)       } |
    id = TK_id TK_dot RW_remove TK_lpar RW_at TK_colon e = exp TK_rpar {$result = instructions.NewRemove($id.line, $id.pos, $id.text, $e.result)} ;

decmatrix returns[interfaces.Instruction result] :
    d = RW_var id = TK_id TK_colon t = typematrix TK_equ df = defmatrix {vecType := $t.result;; $result = instructions.NewInitMatrix($d.line, $d.pos, true,  $id.text, &vecType, $df.result)} |
    d = RW_var id = TK_id TK_equ df = defmatrix                         {$result = instructions.NewInitMatrix($d.line, $d.pos, true,  $id.text, nil, $df.result)} ;

defmatrix returns[*vector.Vector result] :
    m = matrix       {$result = $m.result} |
    s = simplematrix {$result = vector.NewMatrixRepeating($s.result)} ;

matrix returns[*vector.Vector result] :
    TK_lbrk vs = vectors TK_rbrk {$result = vector.NewMatrix(nil, $vs.result)} ;

vectors returns[[]*vector.Vector result] :
    vs = vectors TK_comma v = vector {$result = $vs.result;; $result = append($result, $v.result)} |
    v = vector                       {$result = []*vector.Vector{$v.result}                      } ;

vector returns[*vector.Vector result] :
    TK_lbrk l = listexp TK_rbrk {$result = vector.NewVector(nil, $l.result)} |
    m = matrix                  {$result = $m.result                       } ;

simplematrix returns[*vector.Repeating result] :
    t = typematrix TK_lpar RW_repeating TK_colon r = simplematrix TK_comma RW_count TK_colon c = exp TK_rpar {$result = vector.NewRepeating($t.result.Length, $r.result.Dims + 1, $t.result.Type.Value.(utils.Type), nil, $c.result, $r.result)} |
    t = typematrix TK_lpar RW_repeating TK_colon v = exp TK_comma RW_count TK_colon c = exp TK_rpar          {$result = vector.NewRepeating($t.result.Length, 1, $t.result.Type.Value.(utils.Type), $v.result, $c.result, nil)                 } ;

typematrix returns[utils.VectorType result] :
    TK_lbrk tm = typematrix TK_rbrk {$result = $tm.result;; $result.Length += 1            } |
    b = TK_lbrk t = type TK_rbrk    {$result = utils.VectorType{Line: $b.line, Column: $b.pos, Length: 1, Type: $t.result}} ;

reasignvector returns[interfaces.Instruction result] :
    id = TK_id index = dims TK_equ e = exp {$result = instructions.NewAsignPosArray($id.line, $id.pos, $id.text, $index.result, $e.result)} ;

dims returns[[]interfaces.Expression result] :
    l = dims TK_lbrk e = exp TK_rbrk {$result = $l.result;; $result = append($result, $e.result)} |
    TK_lbrk e = exp TK_rbrk          {$result = []interfaces.Expression{$e.result}              } ;

defstruct returns[interfaces.Instruction result] :
    RW_struct TK_id TK_lbrc listattribs TK_rbrc ;

listattribs :
    listattribs TK_semicolon? attrib |
    attrib TK_semicolon?             ;

attrib :
    (RW_let | RW_var) TK_id (TK_colon typeComp)? (TK_equ exp)? |
    RW_mutating? declfunc ;

decstruct returns[interfaces.Instruction result] :
    (RW_let | RW_var) TK_id (TK_colon TK_id)? TK_equ TK_id TK_lpar listdupla? TK_rpar |
    (RW_let | RW_var) TK_id (TK_colon TK_id)? TK_equ TK_id TK_lpar TK_rpar       ;

listdupla :
    TK_id TK_colon exp TK_comma listdupla |
    TK_id TK_colon exp                    ;

useattribs returns[interfaces.Expression result] :
    obj useattribs1     |
    obj TK_dot callfunc ;

obj :
    TK_id TK_lbrk exp TK_rbrk |
    TK_id                     ;

useattribs1 :
    TK_dot TK_id useattribs1 |
    TK_dot TK_id             ;

print returns[interfaces.Instruction result]:
    p = RW_print TK_lpar exps = listexp TK_rpar {$result = instructions.NewPrint($p.line, $p.pos, $exps.result)} |
    p = RW_print TK_lpar TK_rpar                {$result = instructions.NewPrint($p.line, $p.pos, nil)         } ;

env returns[interfaces.Instruction result] :
    l = TK_lbrc ins = instructions TK_rbrc {$result = instructions.NewBlock($l.line, $l.pos, $ins.result)    } |
    l = TK_lbrc TK_rbrc                    {$result = instructions.NewBlock($l.line, $l.pos, []interface{}{})} ;

instructions returns[[]interface{} result] :
    l = instructions i = instruction {$result = $l.result;; $result = append($result, $i.result)} |
    i = instruction                  {$result = []interface{}{$i.result}                        } ;

instruction returns[interface{} result] :
    inst1 =  decvar                          TK_semicolon? {$result = $inst1.result} |
    inst2 =  deccst                          TK_semicolon? {$result = $inst2.result} |
    inst3 =  ifstruct                                      {$result = $inst3.result} |
    inst4 =  switchstruct                                  {$result = $inst4.result} |
    inst5 =  loopfor                                       {$result = $inst5.result} |
    inst6 =  loopwhile                                     {$result = $inst6.result} |
    inst7 =  guard                                         {$result = $inst7.result} |
    (RW_self TK_dot)? inst8 = reasign        TK_semicolon? {$result = $inst8.result} |
    (RW_self TK_dot)? inst9 = addsub         TK_semicolon? {$result = $inst9.result} |
    inst10 = decvector                       TK_semicolon? {$result = $inst10.result} |
    inst11 = funcvector                      TK_semicolon? {$result = $inst11.result} |
    (RW_self TK_dot)? inst12 = reasignvector TK_semicolon? {$result = $inst12.result} |
    inst13 = decmatrix                       TK_semicolon? {$result = $inst13.result} |
    inst14 = decstruct                       TK_semicolon? {$result = $inst14.result} |
    (RW_self TK_dot)? inst15 = useattribs    TK_semicolon? {$result = $inst15.result} |
    (RW_self TK_dot)? inst16 = callfunc      TK_semicolon? {$result = $inst16.result} |
    inst17 = print                           TK_semicolon? {$result = $inst17.result} |
    inst18 = RW_return e = exp               TK_semicolon? {$result = expressions.NewReturn($inst18.line, $inst18.line, $e.result)}|
    inst19 = RW_return                       TK_semicolon? {$result = expressions.NewReturn($inst19.line, $inst19.line, nil)      }|
    inst20 = RW_continue                     TK_semicolon? {$result = instructions.NewContinue($inst20.line, $inst20.line)        }|
    inst21 = RW_break                        TK_semicolon? {$result = instructions.NewBreak($inst21.line, $inst21.line)           };

type returns[utils.AttribsType result] :
    t = RW_String    {$result = *utils.NewAttribsType($t.line, $t.pos, utils.STRING, true) } |
    t = RW_Int       {$result = *utils.NewAttribsType($t.line, $t.pos, utils.INT, true)    } |
    t = RW_Bool      {$result = *utils.NewAttribsType($t.line, $t.pos, utils.BOOLEAN, true)} |
    t = RW_Character {$result = *utils.NewAttribsType($t.line, $t.pos, utils.CHAR, true)   } |
    t = RW_Float     {$result = *utils.NewAttribsType($t.line, $t.pos, utils.FLOAT, true)  } ;

typeComp returns[utils.AttribsType result] :
    t = type  {$result = $t.result} |
    i = TK_id {$result = *utils.NewAttribsType($i.line, $i.pos, $i.text, false)} ;

exp returns[interfaces.Expression result] :
    // Vector
    id = TK_id index = dims        {$result = expressions.NewAccessArray($id.line, $id.pos, $id.text, $index.result)   } |
    e = exp TK_dot RW_isEmpty      {$result = expressions.NewIsEmpty($e.result.LineN(), $e.result.ColumnN(), $e.result)} |
    e = exp TK_dot RW_count        {$result = expressions.NewCount($e.result.LineN(), $e.result.ColumnN(), $e.result)  } |
    // ARITHMETICS
    s = TK_minus e2 = exp                             {$result = expressions.NewArithmetic($s.line, $s.pos, nil, $s.text, $e2.result)                                } |
    e1 = exp s = (TK_mult | TK_div | TK_mod) e2 = exp {$result = expressions.NewArithmetic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_plus | TK_minus)        e2 = exp {$result = expressions.NewArithmetic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // RELATIONALS
    e1 = exp s = (TK_lessequ | TK_moreequ)   e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_less    | TK_more)      e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_equequ  | TK_notequ)    e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // LOGICS
    s = TK_not e2 = exp           {$result = expressions.NewLogic($s.line, $s.pos, nil, $s.text, $e2.result)                                } |
    e1 = exp s = TK_and e2 = exp  {$result = expressions.NewLogic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = TK_or  e2 = exp  {$result = expressions.NewLogic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // CAST
    t = type TK_lpar e = exp TK_rpar {$result = expressions.NewCast($t.result.Line, $t.result.Column, $t.result.Value.(utils.Type), $e.result)} |
    // ATTRIBUTES STRUCT
    (RW_self TK_dot)? u = useattribs {$result = $u.result} |
    // CALL FUNCTION
    (RW_self TK_dot)? c = callfunc {$result = $c.result} |
    // ACCESS ID
    (RW_self TK_dot)? id = TK_id   {$result = expressions.NewAccessID($id.line, $id.pos, $id.text)             } |
    // NIL
    n = RW_nil                     {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.NIL)    } |
    // PRIMITIVES
    p = TK_string                  {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text[1 : len($p.text) - 1], utils.STRING) } |
    p = TK_char                    {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text[1 : len($p.text) - 1], utils.CHAR)   } |
    p = TK_int                     {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.INT)    } |
    p = TK_float                   {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.FLOAT)  } |
    p = RW_true                    {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.BOOLEAN)} |
    p = RW_false                   {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.BOOLEAN)} |
    // GROUP
    TK_lpar e = exp TK_rpar        {$result = $e.result} ;