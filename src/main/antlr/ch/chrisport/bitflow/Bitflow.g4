/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline+;

fork : NAME? transform_parameters? scheduling_hints? '{' sub_pipeline+'}';

multiinput : '{' (transform ';')+ '}';

pipeline : PIPE_NAME? (multiinput | transform) ((PIPE transform) | (PIPE fork))* ( EOP | EOF );

sub_pipeline : PIPE_NAME? transform ((PIPE transform) | (PIPE fork))* ( EOP | EOF );

// special transform shortcuts:
console: '-';

ipport: IP? PORT;

file: FILEPATH;

full_transform : NAME transform_parameters? scheduling_hints?;

transform : (ipport | console | file | full_transform);

parameter : NAME '=' (STRING | NUMBER | FILEPATH);

transform_parameters : '(' (parameter (',' parameter)*)? ')';

scheduling_hints : '[' (parameter (',' parameter)*)? ']';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

PIPE_NAME : ('"' NAME '":' | NAME ':');

IP: (NUMBER '.' NUMBER '.' NUMBER '.' NUMBER )+;

PORT: ':' NUMBER;

FILEPATH: '"' .?[/\\][-.a-zA-Z0-9:/\\]+ '"';

STRING : '"' .*? '"';

NUMBER : [0-9]+;

NAME : LETTER+;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip;

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? NEWLINE -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;
