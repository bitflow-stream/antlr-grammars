/*
* Parser Rules
*/
grammar Bitflow ;

/*
* Parser Rules
*/

script : multiInputPipeline EOF ;

// Simple tokens
dataInput : name+ schedulingHints? ;
dataOutput : name schedulingHints? ;
name : IDENTIFIER | NUMBER | BOOL | STRING ;
val : NUMBER | BOOL | STRING ;

// Parameters
parameter : name EQ val;
transformParameters : OPEN_PARAMS (parameter (SEP parameter)*)? SEP? CLOSE_PARAMS;

// Pipeline and steps
pipeline : (dataInput | OPEN multiInputPipeline CLOSE) (NEXT pipelineElement)* ;
multiInputPipeline : pipeline (EOP pipeline)* EOP? ;
pipelineElement : transform | fork | multiplexFork | window | dataOutput ;
transform: name transformParameters schedulingHints? ;

// Forks
fork : name transformParameters schedulingHints? OPEN namedSubPipeline (EOP namedSubPipeline)* EOP? CLOSE ;
namedSubPipeline : name+ NEXT subPipeline ;
subPipeline : pipelineElement (NEXT pipelineElement)* ;
multiplexFork : OPEN multiplexSubPipeline (EOP multiplexSubPipeline)* EOP? CLOSE ;
multiplexSubPipeline : subPipeline ;

// Windows
window : WINDOW transformParameters schedulingHints? OPEN windowPipeline CLOSE ;
windowPipeline : transform (NEXT transform)* ;

// Scheduling hints
schedulingHints : OPEN_HINTS (schedulingParameter (SEP schedulingParameter)*)? SEP? CLOSE_HINTS ;
schedulingParameter : parameter ;

/*
* Lexer Rules
*/

OPEN : '{' ;
CLOSE : '}' ;
EOP : ';' ; // End Of Pipeline
NEXT : '->' ;

OPEN_PARAMS : '(' ;
CLOSE_PARAMS : ')' ;
EQ : '=' ;
SEP : ',' ;

OPEN_HINTS : '[' ;
CLOSE_HINTS : ']' ;

WINDOW: 'window' ; // Only "keyword" in the grammar

STRING : '"' .*? '"' | '\'' .*? '\'' | '`' .*? '`' ; // Three types of string delimiter characters for flexibility
NUMBER : [0-9]+('.'[0-9]+)? ;
BOOL : 'true'|'True'|'false'|'False' ;
IDENTIFIER : [a-zA-Z0-9._:\\/-]+ ;

COMMENT : '#' ~('\n'|'\r')* NEWLINE -> skip;
NEWLINE : ('\r' | '\n' | '\r\n') -> skip;
WHITESPACE : (' ' | '\\s') -> skip;
TAB : '\t' -> skip;
