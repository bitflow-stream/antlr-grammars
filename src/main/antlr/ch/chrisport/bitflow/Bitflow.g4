/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline;

outputFork : name? transformParameters? schedulingHints? '{' (output ';')+ '}';

fork : name? transformParameters? schedulingHints? '{' subPipeline+ '}';

window: 'window' transformParameters? schedulingHints? '{' windowSubPipeline '}';

multiinput : input (';' input)* ';'?;

input : name transformParameters? schedulingHints?;

output : name transformParameters? schedulingHints?;

transform: name transformParameters? schedulingHints?;

subPipeline : (CASE? pipelineName PIPE)? (transform | fork | window) (PIPE (transform | fork | window))*  ( EOP )?;

windowSubPipeline : (transform | fork) (PIPE (transform | fork))*  ( EOP )?;

pipeline : (multiinput | input) (PIPE  (transform | fork | window ))* PIPE (output | outputFork) ( EOP | EOF )?;

// special transform shortcuts, removed from script, because they are detected by the bitflow framework itself
// console: '-';
// ipport: IP? PORT;
// file: FILEPATH;
// full_transform: name transform_parameters? scheduling_hints?;
// transform : (ipport | console | file | full_transform);

parameter : NAME '=' (STRING | NUMBER);

transformParameters : '(' (parameter (',' parameter)*)? ')';

name: (STRING | NAME | NUMBER);

pipelineName: (STRING | NAME | NUMBER);

schedulingHints : '[' (parameter (',' parameter)*)? ']';

/*
* Lexer Rules
*/
fragment LETTER : [.a-zA-Z_0-9:\\/-];

fragment LETTER_WITHOUT_DASH : [.a-zA-Z_0-9:\\];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

fragment SINGLE_DASH: '-';

CASE: 'case' | 'CASE';

// End Of Pipeline
EOP : ';';

PIPE : '->';

PIPE_NAME : ('"' NAME '":' | NAME ':');

STRING : '"' .*? '"';

NUMBER : [0-9]+;

NAME : (LETTER*LETTER_WITHOUT_DASH | '"'LETTER*LETTER_WITHOUT_DASH'"' | SINGLE_DASH);

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? NEWLINE -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : (' ' | '\\s') -> skip;

TAB : '\t' -> skip;
