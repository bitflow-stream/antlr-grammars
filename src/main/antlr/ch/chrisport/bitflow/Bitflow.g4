/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline+;

fork : name? transform_parameters? scheduling_hints? '{' sub_pipeline+ '}';

multiinput : '{' (input ';')+ '}';

input : name transform_parameters? scheduling_hints?;

output : name transform_parameters? scheduling_hints?;

sub_pipeline : (CASE? name PIPE)? (fork | transform) ((PIPE transform) | (PIPE fork))*  ( EOP | EOF )?;

pipeline : name? (multiinput | input) ((PIPE transform) | (PIPE fork))* PIPE output ( EOP | EOF )?;

// special transform shortcuts, removed from script, because they are detected by the bitflow framework itself
// console: '-';
// ipport: IP? PORT;
// file: FILEPATH;
// full_transform: name transform_parameters? scheduling_hints?;
// transform : (ipport | console | file | full_transform);
transform: name transform_parameters? scheduling_hints?;

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
