/*
* Parser Rules
*/
grammar Bitflow ;

/*
* Parser Rules
*/

script : pipeline ( EOP pipeline )* EOP? EOF ;

// Simple tokens
input : name+ schedulingHints? ;
output : name schedulingHints? ;
name : IDENTIFIER | NUMBER | BOOL | STRING ;
val : NUMBER | BOOL | STRING ;

// Parameters
parameter : name '=' val;
transformParameters : '(' (parameter (',' parameter)*)? ')';

// Pipeline and steps
pipeline : (input | multiInputPipeline) (NEXT pipelineElement)* ;
multiInputPipeline : '{' pipeline (EOP pipeline)* EOP? '}';
pipelineElement : transform | fork | multiplexFork | window | output ;
transform: name transformParameters schedulingHints? ;

// Forks
fork : name transformParameters schedulingHints? '{' namedSubPipeline (EOP namedSubPipeline)* EOP? '}' ;
namedSubPipeline : name+ NEXT subPipeline ;
subPipeline : pipelineElement (NEXT pipelineElement)* ;
multiplexFork : '{' multiplexSubPipeline (EOP multiplexSubPipeline)* EOP? '}' ;
multiplexSubPipeline : subPipeline ;

// Windows
window : 'window' transformParameters schedulingHints? '{' windowPipeline '}' ;
windowPipeline : transform (NEXT transform)* ;

// Scheduling hints
schedulingHints : '[' (schedulingParameter (',' schedulingParameter)*)? ']' ;
schedulingParameter : parameter ;

/*
* Lexer Rules
*/

EOP : ';' ; // End Of Pipeline
NEXT : '->' ;

STRING : '"' .*? '"' | '\'' .*? '\'' | '`' .*? '`' ;
NUMBER : [0-9]+('.'[0-9]+)? ;
BOOL : 'true'|'True'|'false'|'False' ;
IDENTIFIER : [a-zA-Z0-9._:\\/-]+ ;

COMMENT : '#' ~('\n'|'\r')* NEWLINE -> skip;
NEWLINE : ('\r' | '\n' | '\r\n') -> skip;
WHITESPACE : (' ' | '\\s') -> skip;
TAB : '\t' -> skip;
