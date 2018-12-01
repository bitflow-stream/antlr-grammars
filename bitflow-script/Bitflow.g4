/*
* Parser Rules
*/
grammar Bitflow ;

/*
* Parser Rules
*/

script : pipeline ( EOP pipeline )* EOP? EOF ;

pipeline : (input | multiInputPipeline) (PIPE intermediateTransform)* ;
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
namedSubPipeline : namedSubPipelineKey PIPE subPipeline ;

subPipeline : intermediateTransform (PIPE intermediateTransform)* ;

multiplexFork : '{' multiplexSubPipeline (EOP multiplexSubPipeline)* EOP? '}' ;
multiplexSubPipeline : subPipeline ;
window : 'window' transformParameters schedulingHints? '{' windowPipeline '}' ;
windowPipeline : transform (PIPE transform)* ;

schedulingHints : '[' (parameter (',' parameter)*)? ']' ;

/*
* Lexer Rules
*/

fragment START_LETTER : [a-zA-Z0-9._:\\/-];
fragment END_LETTER : [a-zA-Z0-9._:\\/]; // Identifiers may not end with a dash, to not collide with the '->' operator
fragment SINGLE_DASH: '-';

EOP : ';'; // End Of Pipeline
PIPE : '->';

STRING : '"' .*? '"' | '\'' .*? '\'' | '`' .*? '`';
NUMBER : [0-9]+;
NAME : (START_LETTER (START_LETTER* END_LETTER)?) | SINGLE_DASH;

COMMENT : '#' ~('\n'|'\r')* NEWLINE -> skip;
NEWLINE : ('\r' | '\n' | '\r\n') -> skip;
WHITESPACE : (' ' | '\\s') -> skip;
TAB : '\t' -> skip;
