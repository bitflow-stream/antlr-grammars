/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline+;

fork : name? transform_parameters? scheduling_hints? '{' sub_pipeline+ '}';

multiinput : '{' (input ';')+ '}';

input : (NAME | STRING) transform_parameters? scheduling_hints?;

sub_pipeline : PIPE_NAME? (fork | transform) ((PIPE transform) | (PIPE fork))* ( EOP | EOF )?;

pipeline : PIPE_NAME? (multiinput | input) ((PIPE transform) | (PIPE fork))* ( EOP | EOF );

// special transform shortcuts, removed from script, because they are detected by the bitflow framework itself
// console: '-';
// ipport: IP? PORT;
// file: FILEPATH;
// full_transform: name transform_parameters? scheduling_hints?;
// transform : (ipport | console | file | full_transform);
transform: name transform_parameters? scheduling_hints?;

parameter : NAME '=' (STRING | NUMBER);

transform_parameters : '(' (parameter (',' parameter)*)? ')';

name: (STRING | NAME);

scheduling_hints : '[' (parameter (',' parameter)*)? ']';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9:\\/-];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : (' ' | '\\s') -> skip;

TAB : '\t' -> skip;

// End Of Pipeline
EOP : ';';

PIPE : '->';

PIPE_NAME : ('"' NAME '":' | NAME ':');

STRING : '"' .*? '"';

NUMBER : [0-9]+;

NAME : LETTER+;

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? NEWLINE -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;
