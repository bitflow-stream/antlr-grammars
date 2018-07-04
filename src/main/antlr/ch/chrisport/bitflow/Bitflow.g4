/*
* Parser Rules
*/
grammar Bitflow ;

script : pipeline+;

// special transform shortcuts:
console: '-';

ipport: IP? PORT;

file: FILEPATH;

multiinput : '{' (transform ';')+ '}';

full_transform : NAME transform_parameters? scheduling_hints?;

transform : (ipport | console | file | full_transform);

fork : FORK_NAME? transform_parameters? scheduling_hints? '{' pipeline+'}';

parameter : NAME '=' (STRING | NUMBER | FILEPATH);

parameter_list : (parameter (',' parameter)*)?;

pipeline : PIPE_NAME? (transform | multiinput) ((PIPE transform) | (PIPE fork))* ( EOP | EOF );

transform_parameters : '(' parameter_list ')';

scheduling_hints : '[' parameter_list ']';

/*
* Lexer Rules
*/
fragment LETTER : [a-zA-Z_0-9];

fragment F : ('F'|'f') ;

fragment I : ('I'|'i') ;

// End Of Pipeline
EOP : ';';

PIPE : '->';

FORK_NAME : NAME F 'ork';

PIPE_NAME : ('"' NAME '":' | NAME ':');

IP: (NUMBER '.' NUMBER '.' NUMBER '.' NUMBER )+;

PORT: ':' NUMBER;

FILEPATH: '"' [/\\][-.a-zA-Z0-9:/\\]+ '"';

STRING : '"' .*? '"';

NUMBER : [0-9]+;

NAME : LETTER+;

NEWLINE : ('\r' | '\n') -> skip;

WHITESPACE : ' ' -> skip;

// ignore comments
COMMENT : '//' ~('\n'|'\r')* '\r'? NEWLINE -> skip;

MULTILINE_COMMENT : ('/*' .*? '*/') -> skip ;
