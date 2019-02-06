/*
* Parser Rules
*/
grammar Bitflow ;

script : pipelines EOF ;

// Simple tokens
dataInput : name+ schedulingHints? ;
dataOutput : name schedulingHints? ;
name : IDENTIFIER | STRING ;

// Parameters
parameter : name EQ name ;
parameterList : parameter (SEP parameter)* ;
parameters : OPEN_PARAMS (parameterList SEP?)? CLOSE_PARAMS ;

// Pipelines and steps
pipelines : pipeline (EOP pipeline)* EOP? ;
pipeline : (dataInput | pipelineElement | OPEN pipelines CLOSE) (NEXT pipelineTailElement)* ;
pipelineElement : processingStep | fork | window ;
pipelineTailElement : pipelineElement | multiplexFork | dataOutput ;
processingStep : name parameters schedulingHints? ;

// Forks
fork : name parameters schedulingHints? OPEN namedSubPipeline (EOP namedSubPipeline)* EOP? CLOSE ;
namedSubPipeline : name+ NEXT subPipeline ;
subPipeline : pipelineTailElement (NEXT pipelineTailElement)* ;
multiplexFork : OPEN subPipeline (EOP subPipeline)* EOP? CLOSE ;

// Windows
window : WINDOW parameters schedulingHints? OPEN processingStep (NEXT processingStep)* CLOSE ;

// Scheduling hints
schedulingHints : OPEN_HINTS (parameterList SEP?)? CLOSE_HINTS ;

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
IDENTIFIER : [a-zA-Z0-9._:\\/-]+ ;

COMMENT : '#' ~('\n'|'\r')* NEWLINE -> skip;
NEWLINE : ('\r' | '\n' | '\r\n') -> skip;
WHITESPACE : (' ' | '\\s') -> skip;
TAB : '\t' -> skip;
