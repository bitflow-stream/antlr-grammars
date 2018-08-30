/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline+;

outputFork : name? transform_parameters? scheduling_hints? '{' (output ';')+ '}';

fork : name? transform_parameters? scheduling_hints? '{' sub_pipeline+ '}';

window: 'window' transform_parameters? scheduling_hints? '{' sub_pipeline+ '}';

multiinput : '{' (input ';')+ '}';

input : name transform_parameters? scheduling_hints?;

output : name transform_parameters? scheduling_hints?;

transform: name transform_parameters? scheduling_hints?;

sub_pipeline : (CASE? name PIPE)? (transform | fork | window) (PIPE (transform | fork | window))*  ( EOP | EOF )?;

pipeline : name? (multiinput | input) (PIPE  (transform | fork | window ))* PIPE (output | outputFork) ( EOP | EOF )?;

// special transform shortcuts, removed from script, because they are detected by the bitflow framework itself
// console: '-';
// ipport: IP? PORT;
// file: FILEPATH;
// full_transform: name transform_parameters? scheduling_hints?;
// transform : (ipport | console | file | full_transform);

parameter : NAME '=' (STRING | NUMBER);

transform_parameters : '(' (parameter (',' parameter)*)? ')';

name: (STRING | NAME | NUMBER);

scheduling_hints : '[' (parameter (',' parameter)*)? ']';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9:\\/-];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

CASE: 'case' | 'CASE';

// End Of Pipeline
EOP : ';';

PIPE : '->';

PIPE_NAME : ('"' NAME '":' | NAME ':');

STRING : '"' .*? '"';

NUMBER : [0-9]+;

NAME : (LETTER+ | '"'LETTER+'"');

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? NEWLINE -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : (' ' | '\\s') -> skip;

TAB : '\t' -> skip;
