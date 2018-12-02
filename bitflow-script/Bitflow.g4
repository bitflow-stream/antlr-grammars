/*
* Parser Rules
*/
grammar Bitflow ;

/*
* Parser Rules
*/

script : pipeline ( EOP pipeline )* EOP? EOF ;

pipeline : (input | multiInputPipeline) (NEXT intermediateTransform)* ;
multiInputPipeline : '{' pipeline (EOP pipeline)* EOP? '}';

// TODO support for multiple input declarations (separated by spaces)
// multiinput : input (';' input)* ';'?;

input : endpoint schedulingHints? ;
output : endpoint schedulingHints? ;
name : NAME | STRING ;
namedSubPipelineKey: NAME | STRING | NUMBER ;
endpoint : STRING | NAME ;
val : STRING | NUMBER ; // TODO explicitely add BOOL and FLOAT

parameter : name '=' val;
transformParameters : '(' (parameter (',' parameter)*)? ')';

intermediateTransform : transform | fork | multiplexFork | window | output ;
transform: name transformParameters schedulingHints? ;

fork : name transformParameters schedulingHints? '{' namedSubPipeline (EOP namedSubPipeline)* EOP? '}' ;
namedSubPipeline : namedSubPipelineKey NEXT subPipeline ;

subPipeline : intermediateTransform (NEXT intermediateTransform)* ;

multiplexFork : '{' multiplexSubPipeline (EOP multiplexSubPipeline)* EOP? '}' ;
multiplexSubPipeline : subPipeline ;
window : 'window' transformParameters schedulingHints? '{' windowPipeline '}' ;
windowPipeline : transform (NEXT transform)* ;

schedulingHints : '[' (parameter (',' parameter)*)? ']' ;

/*
* Lexer Rules
*/

EOP : ';'; // End Of Pipeline
NEXT : '->';

STRING : '"' .*? '"' | '\'' .*? '\'' | '`' .*? '`';
NUMBER : [0-9]+;
NAME : [a-zA-Z0-9._:\\/-]+;

COMMENT : '#' ~('\n'|'\r')* NEWLINE -> skip;
NEWLINE : ('\r' | '\n' | '\r\n') -> skip;
WHITESPACE : (' ' | '\\s') -> skip;
TAB : '\t' -> skip;
